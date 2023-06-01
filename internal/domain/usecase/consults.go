package usecase

import (
	"context"
	"encoding/json"
	"math"
	"sort"
	"sync"

	"github.com/nguyenvantuan2391996/be-project/internal/domain/helper"
	"github.com/nguyenvantuan2391996/be-project/internal/domain/model"
	"github.com/nguyenvantuan2391996/be-project/internal/domain/repository"
)

const (
	MaxMax = "max-max" // as big as possible
	MinMax = "min-max" // as small as possible
)

type ConsultDomain struct {
	userRepo        repository.IUserRepositoryInterface
	standardRepo    repository.IStandardRepositoryInterface
	scoreRatingRepo repository.IScoreRatingRepositoryInterface
}

func NewConsultDomain(
	userRepo repository.IUserRepositoryInterface,
	standardRepo repository.IStandardRepositoryInterface,
	scoreRatingRepo repository.IScoreRatingRepositoryInterface,
) *ConsultDomain {
	return &ConsultDomain{
		userRepo:        userRepo,
		standardRepo:    standardRepo,
		scoreRatingRepo: scoreRatingRepo,
	}
}

func (c *ConsultDomain) Consult(ctx context.Context, userId string) ([]*model.ConsultResult, error) {
	var (
		wg               sync.WaitGroup
		weights          []float64
		types            []string
		matrix           [][]float64
		listName         []string
		bestAlternative  []float64
		worstAlternative []float64
		bestDistance     []float64
		worstDistance    []float64
		similarity       []float64
		err              error
	)

	// Step 1: Create an evaluation matrix consisting of m alternatives and n criteria
	// Get list standards
	wg.Add(1)
	go func() {
		defer wg.Done()
		standards, scopedErr := c.standardRepo.GetStandardByQueries(ctx, map[string]interface{}{
			"user_id": userId,
		})
		if scopedErr != nil {
			err = scopedErr
			return
		}
		// Parse weight
		weights, types = ParseWeight(standards)
	}()

	// Get list score ratings
	wg.Add(1)
	go func() {
		defer wg.Done()
		scoreRatings, scopedErr := c.scoreRatingRepo.GetScoreRatingByListQueries(ctx, map[string]interface{}{
			"user_id": userId,
		}, []string{"created_at", "asc"})
		if scopedErr != nil {
			err = scopedErr
			return
		}
		// Parse metadata score rating
		matrix, listName, scopedErr = ParseMetadata(scoreRatings)
		if scopedErr != nil {
			err = scopedErr
			return
		}
	}()

	wg.Wait()
	if err != nil {
		return nil, err
	}

	// Step 2: The matrix is then normalised to form the matrix
	matrix = NormalizeMatrix(matrix)
	weights = NormalizeWeights(weights)

	// Step 3: Calculate the weighted normalised decision matrix
	matrix = CalculateTheWeighted(matrix, weights)

	// Step 4: Determine the worst alternative and the best alternative
	bestAlternative, worstAlternative = GetBestAndWorst(matrix, types)

	// Step 5: Calculate the L2-distance between the target alternative and the worst condition
	bestDistance, worstDistance = CalculateDistance(matrix, bestAlternative, worstAlternative)

	// Step 6: Calculate the similarity to the worst condition
	similarity = CalculateSimilarity(bestDistance, worstDistance)

	return ConvertToListConsultResult(similarity, listName), nil
}

func ParseWeight(standards []*model.Standard) (weights []float64, types []string) {
	for _, value := range standards {
		weights = append(weights, float64(value.Weight))
		types = append(types, value.Type)
	}
	return weights, types
}

func ParseMetadata(scoreRatings []*model.ScoreRating) (matrix [][]float64, listName []string, err error) {
	var metadataStruct []*model.MetadataScoreRating
	for _, value := range scoreRatings {
		err = json.Unmarshal([]byte(value.Metadata), &metadataStruct)
		if err != nil {
			return nil, nil, err
		}
		var score []float64
		for _, v := range metadataStruct {
			score = append(score, v.Score)
		}
		matrix = append(matrix, score)
		if len(metadataStruct) > 0 {
			listName = append(listName, metadataStruct[0].Name)
		}
	}
	return matrix, listName, nil
}

func NormalizeMatrix(matrix [][]float64) [][]float64 {
	for col := 0; col < len(matrix[0]); col++ {
		var sumSquare float64
		for row := 0; row < len(matrix); row++ {
			sumSquare += matrix[row][col] * matrix[row][col]
		}
		// Set data is normalized
		for row := 0; row < len(matrix); row++ {
			matrix[row][col] /= math.Sqrt(sumSquare)
		}
	}
	return matrix
}

func NormalizeWeights(weights []float64) []float64 {
	var sum float64
	for _, value := range weights {
		sum += value
	}
	for index, value := range weights {
		weights[index] = value / sum
	}
	return weights
}

func CalculateTheWeighted(matrix [][]float64, weights []float64) [][]float64 {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			matrix[row][col] *= weights[col]
		}
	}
	return matrix
}

func GetBestAndWorst(matrix [][]float64, types []string) (bestAlternative, worstAlternative []float64) {
	var listColumns [][]float64
	for col := 0; col < len(matrix[0]); col++ {
		var columns []float64
		for row := 0; row < len(matrix); row++ {
			columns = append(columns, matrix[row][col])
		}
		listColumns = append(listColumns, columns)
	}

	for index, value := range listColumns {
		if types[index] == MaxMax {
			bestAlternative = append(bestAlternative, helper.FindMax(value))
			worstAlternative = append(worstAlternative, helper.FindMin(value))
		} else {
			bestAlternative = append(bestAlternative, helper.FindMin(value))
			worstAlternative = append(worstAlternative, helper.FindMax(value))
		}
	}
	return bestAlternative, worstAlternative
}

func CalculateDistance(matrix [][]float64, bestAlternative, worstAlternative []float64) (bestDistance, worstDistance []float64) {
	for _, value := range matrix {
		bestDistance = append(bestDistance, helper.CalculateEuclideanDistance(bestAlternative, value))
		worstDistance = append(worstDistance, helper.CalculateEuclideanDistance(worstAlternative, value))
	}
	return bestDistance, worstDistance
}

func CalculateSimilarity(bestDistance, worstDistance []float64) (similarity []float64) {
	if len(bestDistance) != len(worstDistance) {
		return []float64{}
	}
	for index := 0; index < len(bestDistance); index++ {
		similarity = append(similarity, worstDistance[index]/(bestDistance[index]+worstDistance[index]))
	}
	return similarity
}

func ConvertToListConsultResult(similarity []float64, listName []string) (list []*model.ConsultResult) {
	if len(similarity) != len(listName) {
		return []*model.ConsultResult{}
	}
	for index := 0; index < len(similarity); index++ {
		list = append(list, &model.ConsultResult{
			Name:       listName[index],
			Similarity: similarity[index],
		})
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Similarity > list[j].Similarity
	})
	return list
}

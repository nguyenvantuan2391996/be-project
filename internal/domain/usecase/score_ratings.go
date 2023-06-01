package usecase

import (
	"context"

	"github.com/nguyenvantuan2391996/be-project/handler/model"
	modelDomain "github.com/nguyenvantuan2391996/be-project/internal/domain/model"
	"github.com/nguyenvantuan2391996/be-project/internal/domain/repository"
)

type ScoreRatingDomain struct {
	scoreRatingRepo repository.IScoreRatingRepositoryInterface
}

func NewScoreRatingDomain(
	scoreRatingRepo repository.IScoreRatingRepositoryInterface,
) *ScoreRatingDomain {
	return &ScoreRatingDomain{
		scoreRatingRepo: scoreRatingRepo,
	}
}

func (s *ScoreRatingDomain) BulkCreateScoreRating(ctx context.Context, scoreRatingReq []*model.ScoreRatingRequest) error {
	var scoreRatings []*modelDomain.ScoreRating
	for _, value := range scoreRatingReq {
		v := &modelDomain.ScoreRating{
			UserId:   value.UserId,
			Metadata: value.Metadata,
		}
		if err := v.GenerateID(); err != nil {
			return err
		}
		scoreRatings = append(scoreRatings, v)
	}
	return s.scoreRatingRepo.BulkCreateScoreRating(ctx, scoreRatings)
}

func (s *ScoreRatingDomain) GetListScoreRating(ctx context.Context, scoreRating *model.ScoreRating) ([]*modelDomain.ScoreRating, error) {
	return s.scoreRatingRepo.GetScoreRatingByListQueries(ctx, map[string]interface{}{
		"user_id": scoreRating.UserId,
	}, []string{"created_at", "asc"})
}

func (s *ScoreRatingDomain) UpdateScoreRating(ctx context.Context, scoreRating *model.ScoreRating) error {
	data, err := s.scoreRatingRepo.GetScoreRatingByListQueries(ctx, map[string]interface{}{
		"id": scoreRating.ID,
	}, []string{"created_at", "asc"})
	if err != nil {
		return err
	}

	return s.scoreRatingRepo.UpdateScoreRatingWithMap(ctx, data[0], map[string]interface{}{
		"metadata": scoreRating.Metadata,
	})
}

func (s *ScoreRatingDomain) DeleteScoreRating(ctx context.Context, scoreRating *model.ScoreRating) error {
	return s.scoreRatingRepo.DeleteScoreRatingByQueries(ctx, map[string]interface{}{
		"id": scoreRating.ID,
	})
}

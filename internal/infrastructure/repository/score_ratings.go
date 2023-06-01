package repository

import (
	"context"
	"strings"

	"github.com/nguyenvantuan2391996/be-project/internal/domain/model"
	"gorm.io/gorm"
)

const (
	BatchSizeCreate = 20
)

type ScoreRatingRepository struct {
	db *gorm.DB
}

func NewScoreRatingRepository(db *gorm.DB) *ScoreRatingRepository {
	return &ScoreRatingRepository{db: db}
}

func (sr *ScoreRatingRepository) BulkCreateScoreRating(ctx context.Context, scoreRatings []*model.ScoreRating) error {
	return sr.db.WithContext(ctx).CreateInBatches(&scoreRatings, BatchSizeCreate).Error
}

func (sr *ScoreRatingRepository) GetScoreRatingByListQueries(ctx context.Context, queries map[string]interface{}, sort []string) ([]*model.ScoreRating, error) {
	var result []*model.ScoreRating
	if err := sr.db.WithContext(ctx).Where(queries).Order(strings.Join(sort, " ")).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (sr *ScoreRatingRepository) UpdateScoreRatingWithMap(ctx context.Context, scoreRating *model.ScoreRating, data map[string]interface{}) error {
	if err := sr.db.WithContext(ctx).Model(&scoreRating).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (sr *ScoreRatingRepository) DeleteScoreRatingByQueries(ctx context.Context, queries map[string]interface{}) error {
	return sr.db.WithContext(ctx).Where(queries).Delete(&model.ScoreRating{}).Error
}

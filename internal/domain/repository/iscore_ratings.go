package repository

import (
	"context"

	"github.com/nguyenvantuan2391996/be-project/internal/domain/model"
)

//go:generate mockgen -package=repository -destination=iscore_ratings_mock.go -source=iscore_ratings.go

type IScoreRatingRepositoryInterface interface {
	BulkCreateScoreRating(ctx context.Context, scoreRatings []*model.ScoreRating) error
	GetScoreRatingByListQueries(ctx context.Context, queries map[string]interface{}, sort []string) ([]*model.ScoreRating, error)
	UpdateScoreRatingWithMap(ctx context.Context, scoreRating *model.ScoreRating, data map[string]interface{}) error
	DeleteScoreRatingByQueries(ctx context.Context, queries map[string]interface{}) error
}

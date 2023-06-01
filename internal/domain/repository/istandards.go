package repository

import (
	"context"

	"github.com/nguyenvantuan2391996/be-project/internal/domain/model"
)

//go:generate mockgen -package=repository -destination=istandards_mock.go -source=istandards.go

type IStandardRepositoryInterface interface {
	CreateStandard(ctx context.Context, standard *model.Standard) (*model.Standard, error)
	GetStandardByQueries(ctx context.Context, queries map[string]interface{}) ([]*model.Standard, error)
	DeleteStandardByQueries(ctx context.Context, queries map[string]interface{}) error
}

package repository

import (
	"context"

	"github.com/nguyenvantuan2391996/be-project/internal/domain/model"
)

//go:generate mockgen -package=repository -destination=iusers_mock.go -source=iusers.go

type IUserRepositoryInterface interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
}

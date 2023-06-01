package usecase

import (
	"context"

	"github.com/nguyenvantuan2391996/be-project/internal/domain/model"
	"github.com/nguyenvantuan2391996/be-project/internal/domain/repository"
)

type UserDomain struct {
	userRepo repository.IUserRepositoryInterface
}

func NewUserDomain(
	userRepo repository.IUserRepositoryInterface,
) *UserDomain {
	return &UserDomain{
		userRepo: userRepo,
	}
}

func (u *UserDomain) CreateUser(ctx context.Context, name string) (*model.User, error) {
	return u.userRepo.CreateUser(ctx, &model.User{
		Name: name,
	})
}

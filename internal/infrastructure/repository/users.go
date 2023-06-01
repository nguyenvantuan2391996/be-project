package repository

import (
	"context"

	"github.com/nguyenvantuan2391996/be-project/internal/domain/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	if err := user.Model.GenerateID(); err != nil {
		return nil, err
	}
	if err := u.db.WithContext(ctx).Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

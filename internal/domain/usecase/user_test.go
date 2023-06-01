package usecase

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nguyenvantuan2391996/be-project/internal/domain/model"
	"github.com/nguyenvantuan2391996/be-project/internal/domain/repository"
	"github.com/stretchr/testify/assert"
)

func TestUserDomain_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := repository.NewMockIUserRepositoryInterface(ctrl)
	userRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(&model.User{Name: "Tuan"}, nil)

	userDomain := NewUserDomain(userRepo)
	user, err := userDomain.CreateUser(context.Background(), "Tuan")

	assert.Nil(t, err)
	assert.Equal(t, "Tuan", user.Name)
}

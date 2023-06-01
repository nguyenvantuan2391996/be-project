package usecase

import (
	"context"

	"github.com/nguyenvantuan2391996/be-project/handler/model"
	modelDomain "github.com/nguyenvantuan2391996/be-project/internal/domain/model"
	"github.com/nguyenvantuan2391996/be-project/internal/domain/repository"
)

type StandardDomain struct {
	standardRepo repository.IStandardRepositoryInterface
}

func NewStandardDomain(
	standardRepo repository.IStandardRepositoryInterface,
) *StandardDomain {
	return &StandardDomain{
		standardRepo: standardRepo,
	}
}

func (s *StandardDomain) CreateStandard(ctx context.Context, standardRequest *model.StandardRequest) (*modelDomain.Standard, error) {
	return s.standardRepo.CreateStandard(ctx, &modelDomain.Standard{
		UserId:       standardRequest.UserID,
		StandardName: standardRequest.StandardName,
		Weight:       standardRequest.Weight,
		Type:         standardRequest.Type,
	})
}

func (s *StandardDomain) GetStandards(ctx context.Context, standard *model.Standard) ([]*modelDomain.Standard, error) {
	return s.standardRepo.GetStandardByQueries(ctx, map[string]interface{}{
		"user_id": standard.UserID,
	})
}

func (s *StandardDomain) DeleteStandard(ctx context.Context, standard *model.Standard) error {
	return s.standardRepo.DeleteStandardByQueries(ctx, map[string]interface{}{
		"id": standard.ID,
	})
}

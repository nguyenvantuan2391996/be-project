package repository

import (
	"context"
	"errors"

	"github.com/nguyenvantuan2391996/be-project/internal/domain/model"
	"gorm.io/gorm"
)

type StandardRepository struct {
	db *gorm.DB
}

func NewStandardRepository(db *gorm.DB) *StandardRepository {
	return &StandardRepository{db: db}
}

func (s *StandardRepository) CreateStandard(ctx context.Context, standard *model.Standard) (*model.Standard, error) {
	if err := standard.Model.GenerateID(); err != nil {
		return nil, err
	}
	if err := s.db.WithContext(ctx).Create(&standard).Error; err != nil {
		return nil, err
	}
	return standard, nil
}

func (s *StandardRepository) GetStandardByQueries(ctx context.Context, queries map[string]interface{}) ([]*model.Standard, error) {
	var result []*model.Standard
	if err := s.db.WithContext(ctx).Where(queries).Find(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return result, nil
}
func (s *StandardRepository) DeleteStandardByQueries(ctx context.Context, queries map[string]interface{}) error {
	return s.db.WithContext(ctx).Where(queries).Delete(&model.Standard{}).Error
}

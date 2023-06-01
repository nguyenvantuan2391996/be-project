package repository

import (
	"github.com/nguyenvantuan2391996/be-project/internal/domain/model"
	"gorm.io/gorm"
)

type StatisticalRepository struct {
	db *gorm.DB
}

func NewStatisticalRepository(db *gorm.DB) *StatisticalRepository {
	return &StatisticalRepository{db: db}
}

func (s *StatisticalRepository) Create(record *model.Statistical) error {
	return s.db.Create(record).Error
}

func (s *StatisticalRepository) UpdateWithMap(record *model.Statistical, queries map[string]interface{}) error {
	return s.db.Model(record).Updates(queries).Error
}

func (s *StatisticalRepository) GetByQueries(queries map[string]interface{}) (*model.Statistical, error) {
	var result *model.Statistical
	if err := s.db.Where(queries).First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

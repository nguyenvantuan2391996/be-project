package model

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Model struct {
	ID        string          `gorm:"primary_key;type:varchar(255)" mapstructure:"ignore"`
	CreatedAt time.Time       `gorm:"type:datetime" mapstructure:"ignore"`
	UpdatedAt time.Time       `gorm:"type:datetime" mapstructure:"ignore"`
	DeletedAt *gorm.DeletedAt `gorm:"type:datetime" mapstructure:"ignore"`
}

func (model *Model) GenerateID() error {
	uv4, err := uuid.NewV4()
	if err != nil {
		return err
	}
	model.ID = uv4.String()
	return nil
}

package model

import "github.com/nguyenvantuan2391996/be-project/handler/model"

const (
	MaxMax = "max-max" // as big as possible
	MinMax = "min-max" // as small as possible
)

type Standard struct {
	Model
	UserId       string `json:"user_id"`
	StandardName string `json:"standard_name"`
	Weight       int    `json:"weight"`
	Type         string `json:"type"`
}

func (s *Standard) ToResponse() *model.StandardResponse {
	return &model.StandardResponse{
		ID:           s.ID,
		UserID:       s.UserId,
		StandardName: s.StandardName,
		Weight:       s.Weight,
		Type:         s.convertType(),
	}
}

func (s *Standard) convertType() string {
	if s.Type == MaxMax {
		return "as big as possible"
	}
	if s.Type == MinMax {
		return "as small as possible"
	}

	return s.Type
}

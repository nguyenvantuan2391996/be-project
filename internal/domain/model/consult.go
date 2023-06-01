package model

import "github.com/nguyenvantuan2391996/be-project/handler/model"

type ConsultResult struct {
	Name       string  `json:"name"`
	Similarity float64 `json:"similarity"`
}

type MetadataScoreRating struct {
	Name         string  `json:"name"`
	StandardName string  `json:"standard_name"`
	Score        float64 `json:"score"`
}

func (c *ConsultResult) ToResponse() *model.ConsultResponse {
	return &model.ConsultResponse{
		Name:       c.Name,
		Similarity: c.Similarity,
	}
}

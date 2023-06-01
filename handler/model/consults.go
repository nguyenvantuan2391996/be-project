package model

type ConsultResponse struct {
	Name       string  `json:"name"`
	Similarity float64 `json:"similarity"`
}

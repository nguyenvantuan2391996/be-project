package model

type ScoreRatingRequest struct {
	ID       string `json:"id"`
	UserId   string `json:"user_id" validate:"required"`
	Metadata string `json:"metadata" validate:"required"`
}

type ScoreRatingResponse struct {
	ID       string `json:"id"`
	Metadata string `json:"metadata"`
}

type ScoreRating struct {
	ID       string `json:"id"`
	UserId   string `json:"user_id"`
	Metadata string `json:"metadata"`
}

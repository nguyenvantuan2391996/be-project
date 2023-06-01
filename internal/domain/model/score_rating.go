package model

type ScoreRating struct {
	Model
	UserId   string `json:"user_id"`
	Metadata string `json:"metadata"`
}

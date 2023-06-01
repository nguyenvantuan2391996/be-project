package handler

import (
	"github.com/nguyenvantuan2391996/be-project/internal/domain/usecase"
)

type Handler struct {
	userDomain        *usecase.UserDomain
	standardDomain    *usecase.StandardDomain
	scoreRatingDomain *usecase.ScoreRatingDomain
	consultDomain     *usecase.ConsultDomain
	validate          *validator.Validate
}

func NewHandler(
	userDomain *usecase.UserDomain,
	standardDomain *usecase.StandardDomain,
	scoreRatingDomain *usecase.ScoreRatingDomain,
	consultDomain *usecase.ConsultDomain,
	validate *validator.Validate) *Handler {
	return &Handler{
		userDomain:        userDomain,
		standardDomain:    standardDomain,
		scoreRatingDomain: scoreRatingDomain,
		consultDomain:     consultDomain,
		validate:          validate,
	}
}

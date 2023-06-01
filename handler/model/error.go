package model

type ErrorSystem struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorValidateRequest struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

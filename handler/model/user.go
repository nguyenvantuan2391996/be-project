package model

type UserRequest struct {
	Name string `json:"name" validate:"required,gte=3,lte=20"`
}

type UserResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

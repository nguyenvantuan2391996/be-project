package model

const SUCCESS = "Success"

type DeletedResponse struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

type BulkCreateResponse struct {
	Type   string `json:"type"`
	Status string `json:"status"`
}

type UpdateResponse struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

package models

type ErrorResponse struct {
	Error   string `json:"error"`
	Success bool   `json:"success"`
}

package models

type SuccessResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

package dto

type ResponseDto struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

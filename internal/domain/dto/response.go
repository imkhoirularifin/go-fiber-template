package dto

type ResponseDto struct {
	Message string               `json:"message"`
	Errors  []ErrorValidationDto `json:"errors"`
	Data    any                  `json:"data"`
}

type ErrorValidationDto struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

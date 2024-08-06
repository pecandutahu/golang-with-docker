package dto

type ErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

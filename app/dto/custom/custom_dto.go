package custom

import "net/http"

type Response struct {
	StatusCode int    `json:"status_code" example:"200"`
	Code       int    `json:"code" example:"001"`
	Message    string `json:"message" example:"success"`
}

func NewSuccess() Response {
	return Response{
		StatusCode: http.StatusOK,
		Code:       200,
		Message:    "success",
	}
}

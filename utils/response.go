package utils

import "github.com/gin-gonic/gin"

type SuccessResponse struct {
	Status     int    `json:"status"`
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
	Pagination any    `json:"pagination,omitempty"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"error"`
	Log     string `json:"log"`
}

func SendSuccess(c *gin.Context, status int, message string, data any, pagination any) {
	response := SuccessResponse{
		Status:     status,
		Message:    message,
		Data:       data,
		Pagination: pagination,
	}
	c.JSON(status, response)
}

func SendError(c *gin.Context, status int, message string, log string) {
	response := ErrorResponse{
		Status:  status,
		Message: message,
		Log:     log,
	}
	c.JSON(status, response)
}

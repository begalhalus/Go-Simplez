package helper

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, data interface{}, message string) {
	responData := Response{
		Data:    data,
		Message: message,
		Status:  "200",
	}
	c.JSON(http.StatusOK, responData)
}

func ErrorResponse(c *gin.Context, status int, message string) {
	responData := Response{
		Data:    make(map[string]interface{}),
		Message: message,
		Status:  strconv.Itoa(status),
	}
	c.JSON(status, responData)
}

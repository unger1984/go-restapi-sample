package response

import (
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func NewErrorResponse(c *gin.Context, statusCode int, err error) {
	_ = c.Error(err)
	c.AbortWithStatusJSON(statusCode, errorResponse{err.Error()})
}

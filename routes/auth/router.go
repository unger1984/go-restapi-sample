package auth

import "github.com/gin-gonic/gin"

func NewRoutesFactory(group *gin.RouterGroup) func() {
	return func() {
		group.GET("/test", func(context *gin.Context) {
			context.JSON(200, gin.H{"success": true})
		})
	}
}

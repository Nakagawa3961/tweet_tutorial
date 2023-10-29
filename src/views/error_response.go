package views

import "github.com/gin-gonic/gin"

func ErrorResponse(c *gin.Context, errCode int, err error) {
	c.AbortWithStatusJSON(errCode, gin.H{
		"error_message": err.Error(),
	})
}

package tweet

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Tweet struct {
	Message string `json:"message" binding:"required"`
}

func PostTweet(c *gin.Context) {
	var newTweet Tweet

	if err := c.ShouldBindJSON(&newTweet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Tweet successfully posted", "tweet": newTweet})
}

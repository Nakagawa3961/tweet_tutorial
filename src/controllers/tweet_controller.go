package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryuta06012/tweet_backend/src/models"
	"github.com/ryuta06012/tweet_backend/src/usecases"
)

type TweetController struct {
	TweetInteractor *usecases.TweetInteractor
}

func NewTweetController(interactor *usecases.TweetInteractor) *TweetController {
	return &TweetController{
		TweetInteractor: interactor,
	}
}

type Tweet struct {
	Message string `json:"message" binding:"required"`
}

func (controller *TweetController) PostTweet(c *gin.Context) {
	var newTweet Tweet

	if err := c.ShouldBindJSON(&newTweet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tweetModel := &models.Tweet{
		Message: newTweet.Message,
	}

	if err := controller.TweetInteractor.SaveTweet(tweetModel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save tweet record"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Tweet successfully posted", "tweet": newTweet})
}

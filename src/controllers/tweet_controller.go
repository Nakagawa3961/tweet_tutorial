package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryuta06012/tweet_backend/src/repositories"
	"github.com/ryuta06012/tweet_backend/src/usecases"
	"github.com/ryuta06012/tweet_backend/src/views"
)

type TweetController struct {
	interactor *usecases.TweetInteractor
}

func NewTweetController(mysqlHandler *repositories.MysqlRepository) *TweetController {
	return &TweetController{
		interactor: usecases.NewTweetInteractor(mysqlHandler),
	}
}

func (t *TweetController) PostTweet(c *gin.Context) {
	var request views.TweetRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		views.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}
	if err := t.interactor.CreateTweetsRecord(request); err != nil {
		views.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusOK)
}

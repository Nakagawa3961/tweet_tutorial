package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ryuta06012/tweet_backend/src/repositories"
	"github.com/ryuta06012/tweet_backend/src/tweet"
)

func InitRouter(router *gin.Engine) *gin.Engine {
	mysqlHandler := repositories.NewMysqlHandler()
	userRouter(router, mysqlHandler)

	router.POST("/tweet/", tweet.PostTweet)
	return router
}

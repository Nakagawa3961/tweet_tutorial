package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ryuta06012/tweet_backend/src/controllers"
	"github.com/ryuta06012/tweet_backend/src/repositories"
)

func TweetRouter(r *gin.Engine, mysqlHandler *repositories.MysqlRepository) {
	tweetRouter := r.Group("/tweet")
	controller := controllers.NewTweetController(mysqlHandler)
	{
		tweetRouter.POST("/", controller.PostTweet)
	}
}

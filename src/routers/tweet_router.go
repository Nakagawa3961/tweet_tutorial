package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ryuta06012/tweet_backend/src/controllers"
	"github.com/ryuta06012/tweet_backend/src/repositories"
	"github.com/ryuta06012/tweet_backend/src/usecases"
	"gorm.io/gorm"
)

// Adjust the function signature to accept a *gorm.DB parameter.
func InitializeTweetRoutes(router *gin.Engine, db *gorm.DB) {
	// Pass the gorm.DB instance to NewTweetRepository.
	tweetRepo := repositories.NewTweetRepository(db)
	tweetInteractor := usecases.NewTweetInteractor(tweetRepo)
	tweetController := controllers.NewTweetController(tweetInteractor)

	// Set up the group and route for tweets.
	tweetGroup := router.Group("/tweet")
	{
		tweetGroup.POST("/", tweetController.PostTweet)
	}
}

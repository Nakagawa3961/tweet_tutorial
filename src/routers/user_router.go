package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ryuta06012/tweet_backend/src/controllers"
	"github.com/ryuta06012/tweet_backend/src/repositories"
)

func UserRouter(r *gin.Engine, mysqlHandler *repositories.MysqlRepository) {
	userRouter := r.Group("/user")
	controller := controllers.NewUsersController(mysqlHandler)
	{
		userRouter.POST("/signup", controller.RegisterFrontUser)
	}
}

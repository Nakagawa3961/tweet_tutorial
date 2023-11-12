package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ryuta06012/tweet_backend/src/repositories"
	"gorm.io/gorm"
)

func InitRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	// Initialize the MySQL handler with the existing database connection.
	mysqlHandler := repositories.NewMysqlHandler()

	// Set up user routes.
	UserRouter(router, mysqlHandler)

	// Set up tweet routes.
	InitializeTweetRoutes(router, db)

	return router
}

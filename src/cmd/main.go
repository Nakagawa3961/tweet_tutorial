package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ryuta06012/tweet_backend/src/routers"
)

func main() {
	engine := gin.Default()

	routers.InitRouter(engine)
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	})
	engine.Run(":8080")
}

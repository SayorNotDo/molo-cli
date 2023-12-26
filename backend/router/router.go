package router

import "github.com/gin-gonic/gin"

func Init() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	auth.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return router
}

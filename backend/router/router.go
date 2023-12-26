package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Init() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	auth.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	auth.POST("/login", func(c *gin.Context) {
		var form LoginForm
		if c.ShouldBind(&form) == nil {
			fmt.Printf("username: %v, password: %v\n", form.User, form.Password)
		}
		c.JSON(200, gin.H{
			"message": "success",
		})
	})
	return router
}

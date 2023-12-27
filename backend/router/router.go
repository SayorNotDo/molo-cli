package router

import (
	"fmt"
	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/go-omnibus/proof"
	"molo-cli/backend/pkg/log"
	"time"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Init() *gin.Engine {
	router := gin.Default()

	// cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "DELETE", "PUT", "PATCH", "HEAD"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "Upgrade", "Origin", "Connection", "Accept-Encoding", "Accept-Language", "Host", "x-requested-with"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Use(ginzap.Ginzap(proof.Logger.Z, time.RFC3339, true))

	router.Use(ginzap.RecoveryWithZap(proof.Logger.Z, true))

	auth := router.Group("/auth")
	auth.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	auth.POST("/login", func(c *gin.Context) {
		var form LoginForm
		log.Logger.Info("====================>>>.")
		if c.ShouldBind(&form) == nil {
			fmt.Printf("username: %v, password: %v\n", form.User, form.Password)
		}
		c.JSON(200, gin.H{
			"message": "success",
		})
	})
	return router
}

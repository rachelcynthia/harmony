package router

import (
	"github.com/jmoiron/sqlx"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, db *sqlx.DB) {

	harmonyRouterGroup := router.Group("")
	harmonyRouterGroup.GET("/login")
	harmonyRouterGroup.GET("/register", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
	harmonyRouterGroup.GET("/blogs", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	harmonyRouterGroup.GET("/login1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

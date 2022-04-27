package router

import (
	"harmony/src/controller"
	"harmony/src/repository"
	"harmony/src/service"

	"github.com/jmoiron/sqlx"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, db *sqlx.DB) {
	authRepository := repository.NewAuthRepository(db)

	authService := service.NewAuthService(authRepository)

	authController := controller.NewAuthController(authService)

	harmonyRouterGroup := router.Group("")
	harmonyRouterGroup.GET("/health-check", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "server up",
		})
	})
	harmonyRouterGroup.POST("/login", authController.Login)
	harmonyRouterGroup.POST("/register", authController.Register)
}

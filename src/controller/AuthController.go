package controller

import (
	apiModel "harmony/src/models/api"
	"harmony/src/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
}

// Register implements AuthController
func (c authController) Register(ctx *gin.Context) {
	var req apiModel.Register

	bindingErr := ctx.ShouldBindBodyWith(&req, binding.JSON)
	if bindingErr != nil {
		log.Printf("error in reading request %v", bindingErr)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": bindingErr})
		return
	}

	userDetails, err := c.authService.Register(ctx, req)
	if err != nil {
		log.Printf("error in service %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, userDetails)
}

// Login implements LoginController
func (c authController) Login(ctx *gin.Context) {
	var req apiModel.Login

	bindingErr := ctx.ShouldBindBodyWith(&req, binding.JSON)
	if bindingErr != nil {
		log.Printf("error in reading request %v", bindingErr)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": bindingErr})
		return
	}

	userDetails, err := c.authService.Login(ctx, req)
	if err != nil {
		log.Printf("error in service %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, userDetails)
}

func NewAuthController(authService service.AuthService) AuthController {
	return authController{
		authService: authService,
	}
}

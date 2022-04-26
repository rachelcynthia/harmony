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
}

type authController struct {
	authService service.AuthService
}

// Login implements LoginController
func (c authController) Login(ctx *gin.Context) {
	var req apiModel.Login

	bindingErr := ctx.ShouldBindBodyWith(&req, binding.JSON)
	if bindingErr != nil {
		log.Fatalf("error in reading request %v", bindingErr)
		return
	}

	userDetails, err := c.authService.Login(ctx, req)
	if err != nil {
		log.Fatalf("error in service %v", err)
		return
	}

	ctx.JSON(http.StatusOK, userDetails)

}

func NewAuthController(authService service.AuthService) AuthController {
	return authController{
		authService: authService,
	}
}

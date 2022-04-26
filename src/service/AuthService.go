package service

//go:generate mockgen -source=AuthService.go -destination=../mocks/mock_AuthService.go -package=mocks

import (
	apiModel "harmony/src/models/api"
	"harmony/src/repository"
	"log"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Login(ctx *gin.Context, req apiModel.Login) (apiModel.User, error)
}

type authService struct {
	authRepository repository.AuthRepository
}

func (s authService) Login(ctx *gin.Context, req apiModel.Login) (apiModel.User, error) {
	userDetails, err := s.authRepository.Login(ctx.Request.Context(), req.Email, req.Password)
	if err != nil {
		log.Fatalf("error when accessing db %v", err)
		return apiModel.User{}, err
	}
	return userDetails, nil
}

func NewAuthService(authRepository repository.AuthRepository) AuthService {
	return authService{
		authRepository: authRepository,
	}
}

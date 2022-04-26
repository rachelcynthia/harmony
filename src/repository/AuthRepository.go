package repository

//go:generate mockgen -source=AuthRepository.go -destination=../mocks/mock_AuthRepository.go -package=mocks

import (
	"context"
	apiModel "harmony/src/models/api"
	"harmony/src/models/db"
	"log"

	"github.com/jmoiron/sqlx"
)

type AuthRepository interface {
	Login(ctx context.Context, email string, password string) (apiModel.User, error)
}

type authRepository struct {
	db *sqlx.DB
}

func (r authRepository) Login(ctx context.Context, email string, password string) (apiModel.User, error) {
	var user []db.User

	query := "SELECT NAME, USERNAME, EMAIL FROM USERS WHERE EMAIL = :1 AND PASSWORD = :2"

	err := r.db.SelectContext(ctx, &user, query, email, password)

	if err != nil {
		log.Fatalf("error when retrieving user details %v", err)
		return apiModel.User{}, err
	}
	userDetails := apiModel.User{
		Name:     user[0].Name,
		Username: user[0].Username,
		Email:    user[0].Email,
	}
	return userDetails, nil
}

func NewAuthRepository(db *sqlx.DB) AuthRepository {
	return &authRepository{db: db}
}

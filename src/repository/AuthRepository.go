package repository

//go:generate mockgen -source=AuthRepository.go -destination=../mocks/mock_AuthRepository.go -package=mocks

import (
	"context"
	"errors"
	apiModel "harmony/src/models/api"
	"harmony/src/models/db"
	"log"

	"github.com/jmoiron/sqlx"
)

type AuthRepository interface {
	Login(ctx context.Context, email string, password string) (apiModel.User, error)
	Register(ctx context.Context, req apiModel.Register) error
}

type authRepository struct {
	db *sqlx.DB
}

func (r authRepository) Register(ctx context.Context, req apiModel.Register) error {
	query := "INSERT INTO USERS(NAME, USERNAME, EMAIL, PASSWORD) VALUES ($1, $2, $3, $4)"

	_, err := r.db.ExecContext(ctx, query, req.Name, req.Username, req.Email, req.Password)

	if err != nil {
		log.Printf("error when adding user details %v", err)
		return err
	}
	return nil
}

func (r authRepository) Login(ctx context.Context, email string, password string) (apiModel.User, error) {
	var user []db.User

	query := "SELECT * FROM USERS WHERE EMAIL = $1 AND PASSWORD = $2"

	err := r.db.SelectContext(ctx, &user, query, email, password)

	if err != nil {
		log.Printf("error when retrieving user details %v", err)
		return apiModel.User{}, err
	}

	if len(user) == 0 {
		return apiModel.User{}, errors.New("email/password entered is incorrect")
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

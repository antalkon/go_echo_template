package service

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"backend/internal/models"
	"backend/internal/repository"
	"backend/pkg/hash"
	tokenjwt "backend/pkg/token_jwt"
)

type AuthService struct {
	authRepo *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{authRepo: repo}
}

func (s *AuthService) SignUpClient(u *models.User, c echo.Context) (string, error) {
	pswHsh, err := hash.GenerateHash(u.Password)
	if err != nil {
		return "", err
	}
	u.PasswordHash = pswHsh
	u.ID = uuid.New()
	token, _ := tokenjwt.GenerateJWT(u.ID)

	return token, s.authRepo.CreateUser(u)
}

func (s *AuthService) SignInClient(u *models.User, c echo.Context) (string, error) {
	client, err := s.authRepo.GetUserByPhoneNumber(u.PhoneNumber)
	if err != nil {
		return "", err
	}
	if client == nil {
		return "", errors.New("Not Found")
	}
	if err := hash.ComparePassword(u.Password, client.PasswordHash); err != nil {
		return "", err
	}

	token, _ := tokenjwt.GenerateJWT(client.ID)

	*u = *client

	fmt.Println(token)
	return token, nil
}

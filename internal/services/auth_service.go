package services

import (
	"blog/internal/repositorys"
	"blog/pkg/models"
)

type AuthService struct {
	repository *repositorys.AuthRepository
}

func NewAuthService(authRepo *repositorys.AuthRepository) *AuthService {
	return &AuthService{repository: authRepo}
}

func (s *AuthService) Login(user models.User) error {
	return nil
}

func (s *AuthService) Singup() error {
	return nil
}

func (s *AuthService) Logout() error {
	return nil
}

func (s *AuthService) Refresh() error {
	return nil
}

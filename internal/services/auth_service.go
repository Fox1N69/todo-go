package services

import "blog/internal/repositorys"

type AuthService struct {
	repository *repositorys.AuthRepository
}

func NewAuthService(authRepo *repositorys.AuthRepository) *AuthService {
	return &AuthService{repository: authRepo}
}

package repositorys

import "blog/pkg/models"

type TokenRepository struct {
}

type TokenRepositoryI interface {
	CreateToken(user *models.User, tokenString string) error
}

func NewTokenRepository() *TokenRepository {
	return &TokenRepository{}
}

func (r *TokenRepository) CreateToken(user *models.User, tokenString string) error {
	return nil
}

package repositorys

import (
	"blog/pkg/models"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenRepository struct {
}

type TokenRepositoryI interface {
	CreateToken(user *models.User, tokenString string) error
}

func NewTokenRepository() *TokenRepository {
	return &TokenRepository{}
}

var secretKey = "secret"

func (r *TokenRepository) CreateToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

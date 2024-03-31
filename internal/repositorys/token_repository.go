package repositorys

import (
	"blog/pkg/models"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenRepository struct {
}

type TokenRepositoryI interface {
	GenerateJWTToken(user *models.User) (string, error)
	VerifyToken(tokenString string) (*jwt.Token, error)
}

func NewTokenRepository() *TokenRepository {
	return &TokenRepository{}
}

var secretKey = []byte("secret")

func (r *TokenRepository) GenerateJWTToken(user *models.User) (string, error) {
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

func (r *TokenRepository) VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

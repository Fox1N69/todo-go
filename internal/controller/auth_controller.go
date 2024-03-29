package controller

import (
	"blog/internal/services"
	"blog/pkg/models"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	service *services.AuthService
}

func NewAuthController(serviceAuth *services.AuthService) *AuthController {
	return &AuthController{service: serviceAuth}
}

var secretKey = []byte("secret")

// function generate jwt token
func generateJWTToken(user *models.User) (string, error) {
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

// function verify token
func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// function for hashed password
func hashPassword(password []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	return hash, err
}

// function for password comparison
func checkPasswordHash(password, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}

func (c *AuthController) Login(ctx echo.Context) error {
	return nil
}

func (c *AuthController) SingUp(ctx echo.Context) error {
	return nil
}

func (c *AuthController) Logout(ctx echo.Context) error {
	return nil
}

func (c *AuthController) Refresh(ctx echo.Context) error {
	return nil
}

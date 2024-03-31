package controller

import (
	"blog/internal/services"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	service *services.AuthService
}

func NewAuthController(serviceAuth *services.AuthService) *AuthController {
	return &AuthController{service: serviceAuth}
}

// function verify token

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

func (ac *AuthController) Login(c echo.Context) error {

	return nil
}

func (ac *AuthController) SingUp(c echo.Context) error {
	return nil
}

func (ac *AuthController) Logout(c echo.Context) error {
	return nil
}

func (ac *AuthController) Refresh(c echo.Context) error {
	return nil
}

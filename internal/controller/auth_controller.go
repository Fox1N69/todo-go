package controller

import "golang.org/x/crypto/bcrypt"

type AuthController struct{}

var secretKey = []byte("secret")

func hashPassword(password []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	return hash, err
}

func checkPasswordHash(password, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}

func (c *AuthController) Login() error {
	return nil
}

func (c *AuthController) Register() error {
	return nil
}

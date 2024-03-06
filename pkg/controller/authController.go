package controller

import "golang.org/x/crypto/bcrypt"


func HashPassword(data map[string]string) []byte {
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	
	return password
}

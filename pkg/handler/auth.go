package handler

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)


func (h *Handler) SingUp(c echo.Context) error {
	var data map[string]string

	log.Fatal(c.Bind(&data))

	password, _ := bcrypt.GenerateFromPassword([]byte(data["passwordl"]), 14)

	return nil
}

func (h *Handler) SingIn(c echo.Context) error {
	return nil
}

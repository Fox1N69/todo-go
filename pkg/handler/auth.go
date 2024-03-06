package handler

import (
	"rest/database"
	"rest/models"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	//`"golang.org/x/crypto/bcrypt"
)

func (h *Handler) SingUp(c echo.Context) error {
	var data map[string]string

	if err := c.Bind(&data); err != nil {
		log.Fatal(err)
	}

	newUser := models.Users {
		Username: data["username"],
		Password: []byte(data["password"]),
	}

	database.DB.Create(&newUser)

	return c.JSON(200, newUser)
}

func (h *Handler) SingIn(c echo.Context) error {
	var data map[string]string

	log.Fatal(c.Bind(&data))

	return nil
}

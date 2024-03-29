package handlers

import "blog/internal/services"

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(serviceAuth *services.AuthService) *AuthHandler {
	return &AuthHandler{service: serviceAuth}
}

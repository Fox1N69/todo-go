package middlewares

import "gorm.io/gorm"

type PostMiddleware struct {
	db *gorm.DB
}

type PostMIddlewareI interface {
}

func NewPostMiddleware(db *gorm.DB) *PostMiddleware {
	return &PostMiddleware{db: db}
}

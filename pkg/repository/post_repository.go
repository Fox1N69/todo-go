package repository

import (
	"rest/pkg/models"

	"gorm.io/gorm"
)

type PostRepository struct {
}

type PostRepositoryI interface {
	CreatePost(post models.Posts) error
	GetPostByID(id uint) (*gorm.DB, error)
	UpdatePost(post models.Posts) error
	DeletePost(post models.Posts) error
}

func (pr *PostRepository) CreatePost(post models.Posts) error {
	return nil
}

func (pr *PostRepository) GetPOstByID(post models.Posts) error {
	return nil
}
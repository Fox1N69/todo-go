package repository

import (
	"rest/database"
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

func (pr *PostRepository) Create(post *models.Posts) error {
	return database.DB.Create(post).Error
}

func (pr *PostRepository) GetByID(id uint) (*gorm.DB, error) {
	return nil, nil
}

func (pr *PostRepository) UpdatePost(postToUpdate models.Posts) error {
	return database.DB.Model(&postToUpdate).Updates(&postToUpdate).Error
}

func (pr *PostRepository) DeletePost(post models.Posts) error {
	return nil
}

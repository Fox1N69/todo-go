package repository

import (
	"errors"
	"rest/database"
	"rest/pkg/models"

	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

type PostRepositoryI interface {
	CreatePost(post models.Posts) error
	GetPostByID(id uint) (models.Posts, error)
	UpdatePost(post models.Posts) error
	DeletePost(post models.Posts) error
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (pr *PostRepository) Create(post *models.Posts) error {
	if pr == nil {
		return errors.New("PostREpository is null")
	}
	return database.DB.Create(post).Error
}

func (pr *PostRepository) GetByID(id uint) (*models.Posts, error) {
	var post models.Posts
	if err := database.DB.Where("id = ?", id).First(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (pr *PostRepository) UpdatePost(postToUpdate *models.Posts) error {
	return database.DB.Model(&postToUpdate).Updates(&postToUpdate).Error
}

func (pr *PostRepository) DeletePost(post models.Posts) error {
	return database.DB.Delete(post).Error
}

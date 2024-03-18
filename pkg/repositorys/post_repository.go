package repository

import (
	"errors"
	"rest/pkg/models"

	"gorm.io/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

type PostRepositoryI interface {
	CreatePost(post *models.Posts) error
	GetPostByID(id uint) (models.Posts, error)
	UpdatePost(post *models.Posts) error
	DeletePost(post *models.Posts) error
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{DB: db}
}

func (pr *PostRepository) Create(post *models.Posts) error {
	if pr == nil {
		return errors.New("PostREpository is null")
	}
	return pr.DB.Create(post).Error
}

func (pr *PostRepository) GetByID(id uint) (*models.Posts, error) {
	var post models.Posts
	if err := pr.DB.Where("id = ?", id).First(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (pr *PostRepository) UpdatePost(postToUpdate *models.Posts) error {
	return pr.DB.Model(&postToUpdate).Updates(&postToUpdate).Error
}

func (pr *PostRepository) DeletePost(post *models.Posts) error {
	return pr.DB.Delete(post).Error
}

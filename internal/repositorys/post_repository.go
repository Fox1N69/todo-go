package repositorys

import (
	"blog/pkg/models"

	"gorm.io/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(database *gorm.DB) *PostRepository {
	return &PostRepository{DB: database}
}

func (r PostRepository) GetAll() ([]models.Post, error) {
	var post []models.Post
	if result := r.DB.Find(&post); result.Error != nil {
		return nil, result.Error
	}
	return post, nil
}

func (r *PostRepository) Create(post *models.Post) error {
	return r.DB.Create(&post).Error
}

func (r *PostRepository) GetByID(id uint) (*models.Post, error) {
	var post models.Post
	if err := r.DB.Where("id = ?", id).First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) Update(post *models.Post) error {
	return r.DB.Save(&post).Error
}

func (r *PostRepository) Delete(post *models.Post) error {
	return r.DB.Delete(&post).Error
}

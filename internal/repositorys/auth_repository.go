package repositorys

import (
	"blog/pkg/models"

	"gorm.io/gorm"
)

type AuthRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

func (r *AuthRepository) Create(user *models.User) error {
	return r.DB.Create(&user).Error
}

func (r *AuthRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("id = ?", &user).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *AuthRepository) Update(user *models.User) error {
	return r.DB.Save(&user).Error
}

func (r *AuthRepository) Delete(user *models.User) error {
	return r.DB.Delete(&user).Error
}

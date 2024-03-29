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
	return nil
}

func (r *AuthRepository) GetByID(id uint) ([]models.User, error) {
	return nil, nil
}

func (r *AuthRepository) Update(user *models.User) error {
	return nil
}

func (r *AuthRepository) Delete(user *models.User) error {
	return nil
}

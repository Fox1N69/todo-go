package repositorys

import "gorm.io/gorm"

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(database *gorm.DB) *PostRepository {
	return &PostRepository{DB: database}
}

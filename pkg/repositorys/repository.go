package repository

import "gorm.io/gorm"

type MainRepository struct {
	Post *PostRepository
}

func NewMainRepository(db *gorm.DB) *MainRepository {
	return &MainRepository{
		Post: NewPostRepository(db),
	}
}

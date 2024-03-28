package services

import (
	"blog/internal/repositorys"
	"blog/pkg/models"
)

type PostServeci struct {
	postRepository *repositorys.PostRepository
}

func NewPostServeci(postRepo *repositorys.PostRepository) *PostServeci {
	return &PostServeci{postRepository: postRepo}
}

func (s *PostServeci) GetAllPosts() ([]models.Post, error) {
	return nil, nil
}

func (s *PostServeci) CreatePost(post models.Post) error {
	return nil
}

func (s *PostServeci) UpdatePost(id uint, post models.Post) error {
	return nil
}

func (s *PostServeci) DeletePost(id uint) error {
	return nil
}

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
	posts, err := s.postRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *PostServeci) CreatePost(post *models.Post) error {
	return s.postRepository.Create(post)
}

func (s *PostServeci) UpdatePost(id uint, post *models.Post) error {
	postUpdate, err := s.postRepository.GetByID(id)
	if err != nil {
		return err
	}

	return s.postRepository.Update(postUpdate)
}

func (s *PostServeci) DeletePost(id uint) error {
	post, err := s.postRepository.GetByID(id)
	if err != nil {
		return err
	}

	return s.postRepository.Delete(post)
}

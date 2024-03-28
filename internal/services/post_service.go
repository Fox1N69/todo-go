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
	existingPost, err := s.postRepository.GetByID(id)
	if err != nil {
		return err
	}

	//update message fields based on recaving data
	existingPost.Title = post.Title
	existingPost.Description = post.Description
	existingPost.PostBody = post.PostBody

	err = s.postRepository.Update(existingPost)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostServeci) DeletePost(id uint) error {
	post, err := s.postRepository.GetByID(id)
	if err != nil {
		return err
	}

	return s.postRepository.Delete(post)
}

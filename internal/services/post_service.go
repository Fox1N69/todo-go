package services

import "blog/internal/repositorys"

type PostServeci struct {
	postRepository *repositorys.PostRepository
}

func NewPostServeci(postRepo *repositorys.PostRepository) *PostServeci {
	return &PostServeci{postRepository: postRepo}
}

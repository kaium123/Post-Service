package service

import (
	"post/models"
	"post/repository"
)

type PostServiceInterface interface {
	CreatePost(post *models.Post) ( error)
	ViewPost(postID int)(*models.Post,error)
	UpdatePost(post *models.Post) (error)
}

type PostService struct {
	repository repository.PostRepositoryInterface
}

func NewPostService(repository repository.PostRepositoryInterface) PostServiceInterface {
	return &PostService{repository: repository}
}

func (s *PostService) CreatePost(post *models.Post) error {

	return s.repository.CreatePost(post)
}

func (s *PostService) ViewPost(postID int) (*models.Post,error) {
	return s.repository.ViewPost(postID)
}

func (s *PostService)UpdatePost(post *models.Post) error {
	return s.repository.Update(post)
}
package service

import (
	"post/models"
	"post/repository"
)

type ReactServiceInterface interface {
	CreateReact(React *models.React) error
	ViewReact(ReactID int) (*models.React, error)
	Unlike(React *models.React) error
	Count(postID int) (int, error)
	List(postID int) ([]*int, error)
}

type ReactService struct {
	repository repository.ReactRepositoryInterface
}

func NewReactService(repository repository.ReactRepositoryInterface) ReactServiceInterface {
	return &ReactService{repository: repository}
}

func (s *ReactService) CreateReact(React *models.React) error {
	return s.repository.CreateReact(React)
}

func (s *ReactService) ViewReact(ReactID int) (*models.React, error) {
	return s.repository.ViewReact(ReactID)
}

func (s *ReactService) Unlike(React *models.React) error {
	return s.repository.Unlike(React)
}

func (s *ReactService) Count(postID int) (int, error) {
	return s.repository.Count(postID)
}

func (s *ReactService) List(postID int) ([]*int, error) {
	return s.repository.List(postID)
}

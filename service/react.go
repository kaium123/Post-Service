package service

import (
	"post/models"
	"post/repository"
)

type ReactServiceInterface interface {
	CreateReact(React *models.React) error
	ViewReact(ReactID int) (*models.React, error)
	UpdateReact(React *models.React) error
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

func (s *ReactService) UpdateReact(React *models.React) error {
	return s.repository.Update(React)
}

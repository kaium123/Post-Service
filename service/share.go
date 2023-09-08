package service

import (
	"post/models"
	"post/repository"
)

type ShareServiceInterface interface {
	CreateShare(Share *models.Share) error
	ViewShare(ShareID int) (*models.Share, error)
	UpdateShare(Share *models.Share) error
}

type ShareService struct {
	repository repository.ShareRepositoryInterface
}

func NewShareService(repository repository.ShareRepositoryInterface) ShareServiceInterface {
	return &ShareService{repository: repository}
}

func (s *ShareService) CreateShare(Share *models.Share) error {

	return s.repository.CreateShare(Share)
}

func (s *ShareService) ViewShare(ShareID int) (*models.Share, error) {
	return s.repository.ViewShare(ShareID)
}

func (s *ShareService) UpdateShare(Share *models.Share) error {
	return s.repository.Update(Share)
}

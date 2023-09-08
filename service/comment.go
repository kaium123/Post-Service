package service

import (
	"post/models"
	"post/repository"
)

type CommentServiceInterface interface {
	CreateComment(Comment *models.Comment) ( error)
	ViewComment(CommentID int)(*models.Comment,error)
	UpdateComment(Comment *models.Comment) (error)
}

type CommentService struct {
	repository repository.CommentRepositoryInterface
}

func NewCommentService(repository repository.CommentRepositoryInterface) CommentServiceInterface {
	return &CommentService{repository: repository}
}

func (s *CommentService) CreateComment(Comment *models.Comment) error {

	return s.repository.CreateComment(Comment)
}

func (s *CommentService) ViewComment(CommentID int) (*models.Comment,error) {
	return s.repository.ViewComment(CommentID)
}

func (s *CommentService)UpdateComment(Comment *models.Comment) error {
	return s.repository.Update(Comment)
}
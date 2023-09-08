package service

import (
	"context"
	"post/common/logger"
	"post/models"
	"post/pb"
	"post/repository"
)

type PostServiceInterface interface {
	CreatePost(post *models.Post) (int, error)
	ViewPost(postID int) (*models.Post, error)
	UpdatePost(post *models.Post) error
}

type PostService struct {
	repository repository.PostRepositoryInterface
	gRPCClient pb.AttachmentServiceClient
	redisRepo  repository.RedisRepositoryInterface
}

func NewPostService(gRPCClient pb.AttachmentServiceClient, repository repository.PostRepositoryInterface, redisRepo repository.RedisRepositoryInterface) PostServiceInterface {
	return &PostService{gRPCClient: gRPCClient, repository: repository, redisRepo: redisRepo}
}

func (s *PostService) CreatePost(post *models.Post) (int, error) {

	id, err := s.repository.CreatePost(post)
	requestAttachments := &pb.RequestAttachments{}
	for _, attachment := range post.Attachments {
		tmpAttachment := pb.RequestAttachment{Name: attachment.Name, Path: attachment.Path, SourceType: "post", SourceId: uint64(id)}
		requestAttachments.Attachments = append(requestAttachments.Attachments, &tmpAttachment)
	}

	_, err = s.gRPCClient.CreateMultiple(context.Background(), requestAttachments)
	if err != nil {
		logger.LogError(err)
		return 0, err
	}
	return id, err
}

func (s *PostService) ViewPost(postID int) (*models.Post, error) {
	params := &pb.FindAllRequestParams{SourceId: int64(postID), SourceType: "post"}
	gRPCAttachments, err := s.gRPCClient.FetchAll(context.Background(), params)
	if err != nil {
		return nil, err
	}

	attachments := []models.Attachment{}

	for _, attattachment := range gRPCAttachments.Attachments {
		attachment := models.Attachment{Name: attattachment.Name, Path: attattachment.Path}
		attachments = append(attachments, attachment)
	}

	post, err := s.repository.ViewPost(postID)
	if err != nil {
		return nil, err
	}

	post.Attachments = attachments
	return post, nil
}

func (s *PostService) UpdatePost(post *models.Post) error {
	params := &pb.FindAllRequestParams{SourceId: int64(post.ID), SourceType: "post"}
	gRPCAttachments, err := s.gRPCClient.FetchAll(context.Background(), params)
	if err != nil {
		return err
	}

	ids := &pb.AttachmentIDs{}
	for _, attachment := range gRPCAttachments.Attachments {
		ids.Id = append(ids.Id, int64(attachment.Id))

	}
	ids.SourceId = int64(post.ID)
	ids.SourceType = "post"

	_, err = s.gRPCClient.Delete(context.Background(), ids)
	if err != nil {
		logger.LogError(err)
		return err
	}

	requestAttachments := &pb.RequestAttachments{}
	for _, attachment := range post.Attachments {
		tmpAttachment := pb.RequestAttachment{Name: attachment.Name, Path: attachment.Path, SourceType: "post", SourceId: uint64(post.ID)}
		requestAttachments.Attachments = append(requestAttachments.Attachments, &tmpAttachment)
	}

	_, err = s.gRPCClient.CreateMultiple(context.Background(), requestAttachments)
	if err != nil {
		logger.LogError(err)
		return err
	}

	return s.repository.Update(post)
}

package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"post/common/logger"
	"post/models"
	"post/pb"
	"post/repository"
)

type PostServiceInterface interface {
	CreatePost(post *models.Post) (int, error)
	ViewPost(postID int) (*models.Post, error)
	UpdatePost(post *models.Post) error
	AllPosts(userID int, accessToken string,requestParams models.RequestParams) ([]*models.Post, error)
}

type PostService struct {
	repository  repository.PostRepositoryInterface
	gRPCClient  pb.AttachmentServiceClient
	redisRepo   repository.RedisRepositoryInterface
	commentRepo repository.CommentRepositoryInterface
	likeRepo    repository.ReactRepositoryInterface
}

func NewPostService(
	gRPCClient pb.AttachmentServiceClient,
	repository repository.PostRepositoryInterface,
	commentRepo repository.CommentRepositoryInterface,
	likeRepo repository.ReactRepositoryInterface,
	redisRepo repository.RedisRepositoryInterface) PostServiceInterface {
	return &PostService{gRPCClient: gRPCClient, repository: repository, commentRepo: commentRepo, likeRepo: likeRepo, redisRepo: redisRepo}
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
	comments, err := s.commentRepo.AllComment(postID)
	if err != nil {
		return nil, err
	}

	post.Comments = comments
	reacts, err := s.likeRepo.Count(postID)
	if err != nil {
		return nil, err
	}

	post.Like = reacts
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

func (s *PostService) AllPosts(userID int, accessToken string,requestParams models.RequestParams) ([]*models.Post, error) {
	url := "http://localhost:8089/api/user/view-friends" // Add "http://" as the scheme
	method := "GET"
	byteData, err := sendHttpRequest(accessToken, url, method)
	if err != nil {
		logger.LogError(err)
		return nil, err
	}
	logger.LogInfo(string(byteData))

	friends := []*models.User{}
	err = json.Unmarshal(byteData, &friends)
	if err != nil {
		return nil, err
	}

	ids := []int{userID}
	for _, friend := range friends {
		ids = append(ids, friend.ID)
	}
	posts, err := s.repository.AllPosts(ids,requestParams)

	logger.LogInfo(len(posts))

	attachmentMap := map[int][]models.Attachment{}
	for _, post := range posts {
		params := &pb.FindAllRequestParams{SourceId: int64(post.ID), SourceType: "post"}
		gRPCAttachments, err := s.gRPCClient.FetchAll(context.Background(), params)
		if err != nil {
			return nil, err
		}

		attachments := []models.Attachment{}
		for _, attattachment := range gRPCAttachments.Attachments {
			attachment := models.Attachment{Name: attattachment.Name, Path: attattachment.Path}
			attachments = append(attachments, attachment)
		}
		attachmentMap[post.ID] = attachments

	}
	for i, post := range posts {
		posts[i].Attachments = attachmentMap[post.ID]
		comments, err := s.commentRepo.AllComment(post.ID)
		if err != nil {
			return nil, err
		}
		posts[i].Comments = comments
		reacts, err := s.likeRepo.Count(post.ID)
		if err != nil {
			return nil, err
		}
		posts[i].Like = reacts
	}

	return posts, nil

}

func sendHttpRequest(accessToken string, url string, method string) ([]byte, error) {

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}
	req.Header.Add("Authorization", accessToken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}
	return body, nil
}

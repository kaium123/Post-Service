package repository

import (
	"database/sql"
	"post/common/logger"
	"post/models"
)

type PostRepositoryInterface interface {
	CreatePost(post *models.Post) (int,error)
	ViewPost(postID int) (*models.Post, error)
	Update(post *models.Post) error
}

type PostRepository struct {
	Db     *sql.DB
	logger logger.LoggerInterface
}

func NewPostRepository(Db *sql.DB, logger logger.LoggerInterface) PostRepositoryInterface {
	return &PostRepository{Db: Db, logger: logger}
}

func (r *PostRepository) CreatePost(post *models.Post) (int,error) {
	var lastInsertedID int64

	query := "INSERT INTO posts (content, user_id) VALUES ($1, $2) RETURNING id"

	err := r.Db.QueryRow(query, post.Content, post.UserID).Scan(&lastInsertedID)
	if err != nil {
		return 0, err 
	}

	return int(lastInsertedID), nil
}

func (r *PostRepository) ViewPost(postID int) (*models.Post, error) {
	var post models.Post
	err := r.Db.QueryRow("SELECT * FROM posts WHERE id = $1", postID).Scan(&post.ID, &post.Content, &post.UserID)
	return &post, err
}

func (s *PostRepository) Update(post *models.Post) error {
	_, err := s.Db.Exec("UPDATE posts SET content = $1 WHERE id = $2", post.Content, post.ID)
	return err
}

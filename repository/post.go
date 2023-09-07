package repository

import (
	"database/sql"
	"post/common/logger"
	"post/models"
)

type PostRepositoryInterface interface {
	CreatePost(post *models.Post) error
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

func (r *PostRepository) CreatePost(post *models.Post) error {
	_, err := r.Db.Exec("INSERT INTO posts (content, user_id) VALUES ($1, $2)", post.Content, post.UserID)
	return err
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

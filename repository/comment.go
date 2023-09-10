package repository

import (
	"database/sql"
	"post/common/logger"
	"post/models"
)

type CommentRepositoryInterface interface {
	CreateComment(Comment *models.Comment) error
	ViewComment(CommentID int) (*models.Comment, error)
	Update(Comment *models.Comment) error
	AllComment(postID int) ([]*models.Comment, error)
	Delete(commentID int) error
}

type CommentRepository struct {
	Db     *sql.DB
	logger logger.LoggerInterface
}

func NewCommentRepository(Db *sql.DB, logger logger.LoggerInterface) CommentRepositoryInterface {
	return &CommentRepository{Db: Db, logger: logger}
}

func (r *CommentRepository) CreateComment(Comment *models.Comment) error {
	_, err := r.Db.Exec("INSERT INTO comments (content, post_id) VALUES ($1, $2)", Comment.Content, Comment.PostID)
	return err
}

func (r *CommentRepository) ViewComment(CommentID int) (*models.Comment, error) {
	var Comment models.Comment
	err := r.Db.QueryRow("SELECT id,content,post_id FROM comments WHERE id = $1", CommentID).Scan(&Comment.ID, &Comment.Content, &Comment.PostID)
	return &Comment, err
}

func (s *CommentRepository) Update(Comment *models.Comment) error {
	_, err := s.Db.Exec("UPDATE comments SET content = $1 WHERE id = $2", Comment.Content, Comment.ID)
	return err
}

func (s *CommentRepository) AllComment(postID int) ([]*models.Comment, error) {
	query := "SELECT id,post_id,content FROM comments WHERE post_id = $1"

	rows, err := s.Db.Query(query, postID)
	if err != nil {
		logger.LogError(err)
		return nil, err
	}
	defer rows.Close()

	comments := []*models.Comment{}

	for rows.Next() {
		comment := models.Comment{}
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.Content); err != nil {
			logger.LogError(err)
			return nil, err
		}
		comments = append(comments, &comment)

	}

	if err := rows.Err(); err != nil {
		logger.LogError(err)
		return nil, err
	}
	return comments, nil
}

func (s *CommentRepository) Delete(commentID int) error {
	_, err := s.Db.Exec("DELETE from comments WHERE id = $1", commentID)
	return err
}

package repository

import (
	"database/sql"
	"errors"
	"post/common/logger"
	"post/models"
)

type ReactRepositoryInterface interface {
	CreateReact(React *models.React) error
	ViewReact(ReactID int) (*models.React, error)
	Unlike(React *models.React) error
	Count(postID int) (int, error)
	List(postID int) ([]*int, error)
}

type ReactRepository struct {
	Db     *sql.DB
	logger logger.LoggerInterface
}

func NewReactRepository(Db *sql.DB, logger logger.LoggerInterface) ReactRepositoryInterface {
	return &ReactRepository{Db: Db, logger: logger}
}

func (r *ReactRepository) CreateReact(React *models.React) error {

	query := "SELECT id FROM reacts where post_id = $1 AND post_type = $2 AND reacted_user_id = $3"

	var id int
	err := r.Db.QueryRow(query, React.PostID, React.PostType, React.ReactedUserID).Scan(&id)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return err
	}
	if id != 0 {
		return errors.New("already reacted")
	}

	logger.LogInfo(React.ReactedUserID)
	_, err = r.Db.Exec("INSERT INTO reacts (post_id, reacted_user_id,post_type) VALUES ($1, $2, $3)", React.PostID, React.ReactedUserID, React.PostType)
	if err != nil {
		return err
	}

	return nil
}

func (r *ReactRepository) ViewReact(ReactID int) (*models.React, error) {
	var React models.React
	err := r.Db.QueryRow("SELECT * FROM Reacts WHERE id = $1", ReactID).Scan(&React.ID, &React.PostID, &React.PostID)
	return &React, err
}

func (s *ReactRepository) Unlike(React *models.React) error {
	query := "SELECT id FROM reacts where post_id = $1 AND post_type = $2 AND reacted_user_id = $3"

	var id int
	logger.LogInfo(React)
	err := s.Db.QueryRow(query, React.PostID, React.PostType, React.ReactedUserID).Scan(&id)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return err
	}
	if id == 0 {
		return errors.New("no react")
	}

	_, err = s.Db.Exec("DELETE from reacts WHERE post_id = $1 and reacted_user_id = $2 and post_type = $3", React.PostID, React.ReactedUserID, React.PostType)
	if err!=nil{
		return err
	}
	return err
}

func (s *ReactRepository) Count(postID int) (int, error) {
	query := "SELECT COUNT(*) FROM reacts where post_id = $1"

	var count int
	err := s.Db.QueryRow(query, postID).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *ReactRepository) List(postID int) ([]*int, error) {
	query := "SELECT reacted_user_id FROM reacts WHERE post_id = $1"

	rows, err := s.Db.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ids := []*int{}

	for rows.Next() {
		id := 0
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, &id)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ids, nil
}

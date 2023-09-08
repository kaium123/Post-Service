package repository

import (
	"database/sql"
	"post/common/logger"
	"post/models"
)

type ReactRepositoryInterface interface {
	CreateReact(React *models.React) error
	ViewReact(ReactID int) (*models.React, error)
	Update(React *models.React) error
}

type ReactRepository struct {
	Db     *sql.DB
	logger logger.LoggerInterface
}

func NewReactRepository(Db *sql.DB, logger logger.LoggerInterface) ReactRepositoryInterface {
	return &ReactRepository{Db: Db, logger: logger}
}

func (r *ReactRepository) CreateReact(React *models.React) error {
	_, err := r.Db.Exec("INSERT INTO Reacts (content, user_id) VALUES ($1, $2)", React.ReactType, React.PostID)
	return err
}

func (r *ReactRepository) ViewReact(ReactID int) (*models.React, error) {
	var React models.React
	err := r.Db.QueryRow("SELECT * FROM Reacts WHERE id = $1", ReactID).Scan(&React.ID, &React.ReactType, &React.PostID)
	return &React, err
}

func (s *ReactRepository) Update(React *models.React) error {
	_, err := s.Db.Exec("UPDATE Reacts SET content = $1 WHERE id = $2", React.ReactType, React.ID)
	return err
}

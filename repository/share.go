package repository

import (
	"database/sql"
	"post/common/logger"
	"post/models"
)

type ShareRepositoryInterface interface {
	CreateShare(Share *models.Share) error
	ViewShare(ShareID int) (*models.Share, error)
	Update(Share *models.Share) error
}

type ShareRepository struct {
	Db     *sql.DB
	logger logger.LoggerInterface
}

func NewShareRepository(Db *sql.DB, logger logger.LoggerInterface) ShareRepositoryInterface {
	return &ShareRepository{Db: Db, logger: logger}
}

func (r *ShareRepository) CreateShare(Share *models.Share) error {
	_, err := r.Db.Exec("INSERT INTO Shares (content, user_id) VALUES ($1, $2)", Share.PostID, Share.SharedID)
	return err
}

func (r *ShareRepository) ViewShare(ShareID int) (*models.Share, error) {
	var Share models.Share
	err := r.Db.QueryRow("SELECT * FROM Shares WHERE id = $1", ShareID).Scan(&Share.ID, &Share.PostID, &Share.SharedID)
	return &Share, err
}

func (s *ShareRepository) Update(Share *models.Share) error {
	_, err := s.Db.Exec("UPDATE Shares SET content = $1 WHERE id = $2", Share.SharedID, Share.ID)
	return err
}

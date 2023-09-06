package repository

import (
	"auth/common/logger"
	"auth/models"
	"database/sql"
)

type UserRepositoryInterface interface {
	Register(user models.User) error
}

type UserRepository struct {
	Db     *sql.DB
	logger logger.LoggerInterface
}

func NewUserRepository(Db *sql.DB, logger logger.LoggerInterface) UserRepositoryInterface {
	return &UserRepository{Db: Db, logger: logger}
}

func (r *UserRepository) Register(user models.User) error {
	_, err := r.Db.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

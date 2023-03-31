package repository

import (
	"database/sql"
	"golang/internal/models"
	"golang/internal/repository/mock"
)

type users interface {
	Create(name, email, password string) (int64, error)
	Get(name, password string) (*models.User, error)
	Update(userId int64, field, value string) error
	Delete(id int64) error
}
type Repository struct {
	UserRepo users
}

func NewRepository(db *sql.DB) *Repository {
	repo := &Repository{}

	if db == nil {
		repo.UserRepo = mock.NewRepo()
	}
}

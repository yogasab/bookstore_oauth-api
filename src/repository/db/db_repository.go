package db

import (
	"errors"

	"github.com/bookstore_oauth-api/src/domain/access_token"
)

type DBRepository interface {
	GetByID(string) (*access_token.AccessToken, error)
}

type dbRepository struct {
}

func NewDBRepository() DBRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetByID(string) (*access_token.AccessToken, error) {
	return nil, errors.New("database connection not implementation yet")
}

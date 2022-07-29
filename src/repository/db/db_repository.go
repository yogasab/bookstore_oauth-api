package db

import (
	"errors"

	"github.com/bookstore_oauth-api/src/clients/cassandra"
	"github.com/bookstore_oauth-api/src/domain/access_token"
	"github.com/gocql/gocql"
)

type DBRepository interface {
	GetByID(ID string) (*access_token.AccessToken, error)
	Create(access_token access_token.AccessToken) error
}

type dbRepository struct {
}

func NewDBRepository() DBRepository {
	return &dbRepository{}
}

const (
	queryGetAccessToken    = "SELECT access_token, client_id, expires, user_id FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens (access_token, client_id, expires, user_id) VALUES (?, ?, ?, ?);"
)

func (r *dbRepository) GetByID(ID string) (*access_token.AccessToken, error) {
	var result access_token.AccessToken

	if err := cassandra.GetSession().
		Query(queryGetAccessToken, ID).
		Scan(&result.AccessToken,
			&result.ClientID,
			&result.Expires,
			&result.UserID,
		); err != nil {
		if err == gocql.ErrNotFound {
			return &result, errors.New("no access token found with correspond ID")
		}
		return &result, errors.New("error when trying to get current id")
	}

	return &result, nil
}

func (r *dbRepository) Create(access_token access_token.AccessToken) error {
	if err := cassandra.GetSession().Query(queryCreateAccessToken,
		access_token.AccessToken,
		access_token.ClientID,
		access_token.Expires,
		access_token.UserID,
	).Exec(); err != nil {
		return err
	}
	return nil
}

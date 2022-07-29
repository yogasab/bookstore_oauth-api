package db

import (
	"errors"

	"github.com/bookstore_oauth-api/src/clients/cassandra"
	"github.com/bookstore_oauth-api/src/domain/access_token"
	"github.com/gocql/gocql"
)

type DBRepository interface {
	GetByID(ID string) (*access_token.AccessToken, error)
}

type dbRepository struct {
}

func NewDBRepository() DBRepository {
	return &dbRepository{}
}

const (
	queryGetByID = "SELECT access_token, client_id, expires, user_id FROM access_tokens WHERE access_token=?;"
)

func (r *dbRepository) GetByID(ID string) (*access_token.AccessToken, error) {
	var result access_token.AccessToken

	if err := cassandra.GetSession().
		Query(queryGetByID, ID).
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

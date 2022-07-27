package access_token

import (
	"errors"
	"strings"
)

type Service interface {
	GetByID(accessTokenID string) (*AccessToken, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetByID(accessTokenID string) (*AccessToken, error) {
	accessTokenID = strings.TrimSpace(accessTokenID)

	if len(accessTokenID) == 0 {
		return nil, errors.New("Invalid access token ID")
	}

	accessToken, err := s.repository.GetByID(accessTokenID)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

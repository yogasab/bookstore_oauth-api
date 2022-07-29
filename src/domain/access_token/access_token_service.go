package access_token

import (
	"errors"
	"strings"
)

type Service interface {
	GetAccessTokenByID(accessTokenID string) (*AccessToken, error)
	CreateAccessToken(input CreateAccessTokenInput) (*AccessToken, error)
	UpdatExpiredAccessToken(accessToken UpdateAccessTokenInput) (*AccessToken, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetAccessTokenByID(accessTokenID string) (*AccessToken, error) {
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

func (s *service) CreateAccessToken(input CreateAccessTokenInput) (*AccessToken, error) {
	access_token := AccessToken{}
	access_token.AccessToken = input.AccessToken
	access_token.UserID = input.UserID
	access_token.Expires = input.Expires
	access_token.ClientID = input.ClientID

	if err := s.repository.Create(access_token); err != nil {
		return &access_token, err
	}

	return &access_token, nil
}

func (s *service) UpdatExpiredAccessToken(accessToken UpdateAccessTokenInput) (*AccessToken, error) {
	access_token := AccessToken{}
	access_token.AccessToken = accessToken.AccessToken
	access_token.Expires = accessToken.Expires

	if err := s.repository.UpdateAccessToken(access_token); err != nil {
		return &access_token, err
	}

	return &access_token, nil
}

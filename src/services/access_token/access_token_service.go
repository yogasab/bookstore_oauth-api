package access_token

import (
	"errors"
	"strings"

	"github.com/bookstore_oauth-api/src/domain/access_token"
	"github.com/bookstore_oauth-api/src/domain/user"

	"github.com/bookstore_oauth-api/src/repository/db"
	"github.com/bookstore_oauth-api/src/repository/rest"
)

type Service interface {
	GetAccessTokenByID(accessTokenID string) (*access_token.AccessToken, error)
	CreateAccessToken(input access_token.AccessTokenInput) (*access_token.AccessToken, error)
	UpdatExpiredAccessToken(accessToken access_token.UpdateAccessTokenInput) (*access_token.AccessToken, error)
}

type service struct {
	userRestRepository rest.RestUserRepository
	repository         db.DBRepository
}

func NewService(repository db.DBRepository, userRestRepository rest.RestUserRepository) Service {
	return &service{repository: repository, userRestRepository: userRestRepository}
}

func (s *service) GetAccessTokenByID(accessTokenID string) (*access_token.AccessToken, error) {
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

func (s *service) CreateAccessToken(input access_token.AccessTokenInput) (*access_token.AccessToken, error) {
	if err := input.ValidateInput(); err != nil {
		return nil, err
	}

	inputLogin := user.UserLoginInput{}
	inputLogin.Email = input.Email
	inputLogin.Password = input.Password

	data, err := s.userRestRepository.LoginUser(inputLogin)
	if err != nil {
		return nil, errors.New("User is not registered, please register first")
	}
	user := data.Data.(map[string]interface{})["user"].(map[string]interface{})
	userID := float64(user["id"].(float64))
	UserID := int64(userID)

	at := access_token.AccessToken{}

	newAccessToken := at.GetNewAccessToken(UserID)
	newAccessToken.Generate()

	if err := s.repository.Create(newAccessToken); err != nil {
		return nil, err
	}

	return &newAccessToken, nil
}

func (s *service) UpdatExpiredAccessToken(accessToken access_token.UpdateAccessTokenInput) (*access_token.AccessToken, error) {
	at := access_token.AccessToken{}
	at.AccessToken = accessToken.AccessToken
	at.Expires = accessToken.Expires

	if err := s.repository.UpdateAccessToken(at); err != nil {
		return &at, err
	}

	return &at, nil
}

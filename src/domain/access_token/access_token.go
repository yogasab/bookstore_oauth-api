package access_token

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/bookstore_oauth-api/src/utils/crypto_utils"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grandTypeClientCredentials = "client_credentials"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

type AccessTokenInput struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// Used for password grant type
	Email    string `json:"email"`
	Password string `json:"password"`

	// Used for client_credentials grant type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type CreateAccessTokenInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UpdateAccessTokenInput struct {
	AccessToken string `json:"access_token" binding:"required"`
	Expires     int64  `json:"expires" binding:"required"`
}

func (at AccessToken) GetNewAccessToken(userID int64) AccessToken {
	return AccessToken{
		UserID:  userID,
		Expires: time.Now().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserID, at.Expires))
}

func (at *AccessToken) Validate() error {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.New("invalid access token id")
	}
	if at.UserID <= 0 {
		return errors.New("invalid user id")
	}
	if at.ClientID <= 0 {
		return errors.New("invalid client id")
	}
	if at.Expires <= 0 {
		return errors.New("invalid expiration time")
	}
	return nil
}

func (at *AccessTokenInput) Validate() error {
	switch at.GrantType {
	case grantTypePassword:
		break

	case grandTypeClientCredentials:
		break

	default:
		return errors.New("invalid grant_type parameter")
	}

	return nil
}

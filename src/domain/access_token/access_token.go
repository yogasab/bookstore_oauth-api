package access_token

import "time"

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

type CreateAccessTokenInput struct {
	AccessToken string `json:"access_token" binding:"required"`
	UserID      int64  `json:"user_id" binding:"required"`
	ClientID    int64  `json:"client_id" binding:"required"`
	Expires     int64  `json:"expires" binding:"required"`
}

type UpdateAccessTokenInput struct {
	AccessToken string `json:"access_token" binding:"required"`
	Expires     int64  `json:"expires" binding:"required"`
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

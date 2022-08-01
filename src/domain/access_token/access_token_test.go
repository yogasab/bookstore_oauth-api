package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetAccessExpirationTokenTime(t *testing.T) {
	assert.Equal(t, 24, expirationTime, "expiration time should be 24 hours")
}

// func TestGetAccessToken(t *testing.T) {
// 	at := GetNewAccessToken(1)

// 	assert.False(t, at.IsExpired(), "brand new access token should not be expired")
// 	assert.EqualValues(t, "", at.AccessToken, "brand new access token should not have defined access token id")
// 	assert.True(t, at.UserID == 0, "brand new access token should not have an associated user id")
// }

func TestAccessTokenExpiredTime(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token expiring in three hours should not be expired by now")
}

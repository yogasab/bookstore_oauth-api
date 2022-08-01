package rest

import (
	"net/http"
	"os"
	"testing"

	"github.com/bookstore_oauth-api/src/domain/user"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

var (
	client = resty.New()
)

func TestUserLoginFromAPI(t *testing.T) {
	input := user.UserLoginInput{
		Email:    "ajax@amsterdam.com",
		Password: "password",
	}

	user, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{"email": input.Email, "password": input.Password}).
		Post("http://localhost:5001/api/v1/users/login")

	assert.Nil(t, err)
	assert.NotNil(t, user)
}

// Disable server to running this test
func TestUserLoginTimeoutFromAPI(t *testing.T) {
	input := user.UserLoginInput{
		Email:    "the-email@email.com",
		Password: "the-password",
	}

	response, _ := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{"email": input.Email, "password": input.Password}).
		Post("http://localhost:5001/api/v1/users/login")

	repository := NewRestUserRepository()

	data, errData := repository.LoginUser(input)

	assert.Nil(t, data)
	assert.NotNil(t, errData)
	assert.NotNil(t, response)
	assert.EqualValues(t, http.StatusInternalServerError, errData.Code)
	assert.EqualValues(t, "invalid restclient response when trying to login user", errData.Message)
	assert.EqualValues(t, "failed", errData.Error)
}

func TestUserLoginInvalidCredentials(t *testing.T) {
	input := user.UserLoginInput{
		Email:    "ajax@amsterdam.com",
		Password: "wrongpasssword",
	}

	response, errResponse := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{"email": input.Email, "password": input.Password}).
		Post("http://localhost:5001/api/v1/users/login")

	repository := NewRestUserRepository()
	data, errData := repository.LoginUser(input)

	assert.Nil(t, data)
	assert.Nil(t, errResponse)
	assert.NotNil(t, errData)
	assert.NotNil(t, response)
	assert.EqualValues(t, http.StatusUnauthorized, errData.Code)
	assert.NotNil(t, "invalid credentials", errData.Message)
}

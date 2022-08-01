package repository

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

func TestUserLoginTimeoutFromAPI(t *testing.T) {
	input := user.UserLoginInput{
		Email:    "the-email@email.com",
		Password: "the-password",
	}

	response, errResponse := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{"email": "the-email@email.com", "password": "the-password"}).
		Post("http://localhost:5001/api/v1/users/login")

	repository := NewRestUserRepository()

	data, errData := repository.LoginUser(input)

	assert.Nil(t, errResponse)
	assert.Nil(t, data)
	assert.NotNil(t, response)
	assert.NotNil(t, errData)
	assert.EqualValues(t, http.StatusInternalServerError, errData.Code)
	assert.EqualValues(t, "User is not registered, please register first", errData.Message)
	assert.EqualValues(t, "failed", errData.Error)
}

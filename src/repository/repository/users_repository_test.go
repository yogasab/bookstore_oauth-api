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
		SetBody(`{"email":"ajax@amsterdam.com", "password":"password"}`).
		Post("http://localhost:5001/api/v1/users/login")

	repository := NewRestUserRepository()

	data, errRepo := repository.LoginUser(input)
	assert.Nil(t, err)
	assert.Nil(t, errRepo)
	assert.NotNil(t, user)
	assert.NotNil(t, data)
	assert.EqualValues(t, http.StatusOK, data.Meta.Code)
	assert.EqualValues(t, "Users logged in successfully", data.Meta.Message)
	assert.EqualValues(t, "success", data.Meta.Status)
}

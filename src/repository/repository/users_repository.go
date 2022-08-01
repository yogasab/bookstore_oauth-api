package repository

import (
	"encoding/json"
	"net/http"

	"github.com/bookstore_oauth-api/src/domain/user"
	"github.com/go-resty/resty/v2"
)

type RestUserRepository interface {
	LoginUser(input user.UserLoginInput) (*user.Response, *user.ResponseError)
}

type restUserRepository struct {
}

var (
	usersRestClient = resty.New()
)

func NewRestUserRepository() RestUserRepository {
	return &restUserRepository{}
}

func (r *restUserRepository) LoginUser(input user.UserLoginInput) (*user.Response, *user.ResponseError) {
	var userFormatter user.Response

	response, err := usersRestClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{"email": input.Email, "password": input.Password}).
		Post("http://localhost:5001/api/v1/users/login")

	if response == nil || response.RawResponse == nil {
		// return nil, errors.New("invalid restclient response when trying to login user")
		return nil, user.FormatError(http.StatusInternalServerError, "failed", err.Error())
	}

	if response.StatusCode() > 299 {
		var err user.ResponseError
		if err := json.Unmarshal(response.Body(), &err); err != nil {
			// return nil, errors.New("invalid error interface when trying to login user")
			return nil, user.FormatError(http.StatusInternalServerError, "failed", "User is not registered, please register first")
		}
		return nil, user.FormatError(http.StatusInternalServerError, "failed", "User is not registered, please register first")
	}

	if err := json.Unmarshal(response.Body(), &userFormatter); err != nil {
		// return nil, errors.New("error when trying to unmarshal users login response")
		return nil, user.FormatError(http.StatusInternalServerError, "failed", err.Error())
	}

	return &userFormatter, nil
}

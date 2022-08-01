package rest

import (
	"encoding/json"
	"net/http"

	"github.com/bookstore_oauth-api/src/domain/user"
	"github.com/go-resty/resty/v2"
)

type RestUserRepository interface {
	LoginUser(input user.UserLoginInput) (*user.Response, *user.ResponseError)
}

type restUserRepository struct{}

var (
	usersRestClient = resty.New()
	resultError     = user.ResponseError{
		Code:    http.StatusNotFound,
		Error:   "failed",
		Message: "User not found",
	}
)

func NewRestUserRepository() RestUserRepository {
	return &restUserRepository{}
}

func (r *restUserRepository) LoginUser(input user.UserLoginInput) (*user.Response, *user.ResponseError) {
	var userFormatter user.Response

	response, _ := usersRestClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{"email": input.Email, "password": input.Password}).
		Post("http://localhost:5001/api/v1/users/login")

	if response == nil || response.RawResponse == nil {
		return nil, user.FormatError(http.StatusInternalServerError, "failed", "invalid restclient response when trying to login user")
	}

	if response.StatusCode() > 299 {
		var errResp user.ResponseError
		err := json.Unmarshal(response.Body(), &errResp)
		if err != nil {
			return nil, user.FormatError(http.StatusInternalServerError, "failed", "invalid error interface when trying to login user")
		}
		return nil, &errResp
	}

	if err := json.Unmarshal(response.Body(), &userFormatter); err != nil {
		return nil, user.FormatError(http.StatusInternalServerError, "failed", "error when trying to unmarshal users login response")
	}

	return &userFormatter, nil
}

package repository

import (
	"encoding/json"
	"errors"

	"github.com/bookstore_oauth-api/src/domain/user"
	"github.com/go-resty/resty/v2"
)

type RestUserRepository interface {
	LoginUser(input user.UserLoginInput) (*user.Response, error)
}

type restUserRepository struct {
}

var (
	// usersRestClient = rest.RequestBuilder{
	// 	BaseURL: "http://localhost:5001",
	// 	Timeout: 100 * time.Millisecond,
	// }
	usersRestClient = resty.New()
)

func NewRestUserRepository() RestUserRepository {
	return &restUserRepository{}
}

func (r *restUserRepository) LoginUser(input user.UserLoginInput) (*user.Response, error) {
	response, err := usersRestClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"email":"ajax@amsterdam.com", "password":"password"}`).
		Post("http://localhost:5001/api/v1/users/login")

	if err != nil {
		return nil, err
	}

	if response == nil || response.RawResponse == nil {
		return nil, errors.New("invalid restclient response when trying to login user")
	}

	if response.StatusCode() > 299 {
		var err error
		if err = json.Unmarshal(response.Body(), &err); err != nil {
			return nil, errors.New("invalid error interface when trying to login user")
		}
	}

	var userFormatter user.Response
	if err := json.Unmarshal(response.Body(), &userFormatter); err != nil {
		return nil, errors.New("invalid error interface when trying to login user")
	}

	return &userFormatter, nil
}

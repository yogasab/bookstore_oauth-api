package access_token

type Repository interface {
	GetByID(ID string) (*AccessToken, error)
	Create(access_token AccessToken) error
	UpdateAccessToken(access_token AccessToken) error
}

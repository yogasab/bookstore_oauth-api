package access_token

type Repository interface {
	GetByID(string) (*AccessToken, error)
}

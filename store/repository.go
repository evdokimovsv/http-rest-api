package store

import "github.com/evdokimovsv/http-rest-api.git/model"

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}

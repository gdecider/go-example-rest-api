package store

import "go-rest-api/internal/model"

type UserRepository interface {
	Create(*model.User) error
	Find(int64) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}

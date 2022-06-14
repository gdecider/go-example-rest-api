package teststore

import (
	"go-rest-api/internal/model"
	"go-rest-api/internal/store"
)

type Store struct {
	userRepository *UserRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			store: s,
			users: make(map[int64]*model.User),
		}
	}

	return s.userRepository
}

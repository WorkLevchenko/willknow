package teststore

import (
	"github.com/WorkLevchenko/willknow/internal/app/model"
	"github.com/WorkLevchenko/willknow/internal/app/store"
)

type Store struct {
	UserRepository *UserRepository
}

func New() *Store {
	return &Store{}
}

// С Помощью этого метода "внешний мир" сможет использовать репозиторий User.
func (s *Store) User() store.UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}

	s.UserRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.UserRepository
}

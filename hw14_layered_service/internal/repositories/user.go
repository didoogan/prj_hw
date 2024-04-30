package store

import (
	"hw14/internal/entities"
	"sync"
)

type UserMemoryRepository struct {
	mx    sync.Mutex
	users []*entities.UserWithPassword
}

func NewUserMemoryRepository() *UserMemoryRepository {
	return &UserMemoryRepository{users: make([]*entities.UserWithPassword, 0)}
}

func (r *UserMemoryRepository) Save(u *entities.UserWithPassword) error {
	r.mx.Lock()
	defer r.mx.Unlock()

	r.users = append(r.users, u)
	return nil
}

func (r *UserMemoryRepository) List() ([]*entities.User, error) {
	r.mx.Lock()
	defer r.mx.Unlock()

	users := make([]*entities.User, 0)

	for _, u := range r.users {
		user := &entities.User{Login: u.Login}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserMemoryRepository) Get(u *entities.User) (*entities.UserWithPassword, error) {
	r.mx.Lock()
	defer r.mx.Unlock()

	var dbUser *entities.UserWithPassword

	for _, user := range r.users {
		if user.Login == u.Login {
			dbUser = user
			break
		}
	}
	return dbUser, nil
}

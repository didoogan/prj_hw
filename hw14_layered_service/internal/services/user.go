package services

import (
	"hw14/internal/entities"
	store "hw14/internal/repositories"
)

type User struct {
	repository store.UserRepository
}

func NewUserService(r store.UserRepository) *User {
	return &User{repository: r}
}

func (u *User) List() ([]*entities.User, error) {
	return u.repository.List()
}

func (u *User) Save(user *entities.UserWithPassword) (*entities.User, error) {
	err := u.repository.Save(user)
	if err != nil {
		return nil, err
	}

	return &entities.User{Login: user.Login}, nil
}

func (u *User) CheckPassword(user *entities.UserWithPassword) (bool, error) {
	dbUser, err := u.repository.Get(&entities.User{Login: user.Login})
	if err != nil {
		return false, err
	}

	if dbUser != nil {
		return dbUser.Password == user.Password, nil
	}
	return false, nil
}

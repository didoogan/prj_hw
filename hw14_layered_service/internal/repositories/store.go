package store

import "hw14/internal/entities"

type UserRepository interface {
	Save(*entities.UserWithPassword) error
	List() ([]*entities.User, error)
	Get(*entities.User) (*entities.UserWithPassword, error)
}

type CacheRepository interface {
	Get(key string) (value string, exists bool, error error)
	Set(key, value string, ttl int) error
}

type Store struct {
	User  UserRepository
	Cache CacheRepository
}

package store

import (
	"errors"
	"hw10/entities"
	"sync"
)

type MemStore struct {
	mx    sync.Mutex
	tasks []*entities.Task
}

var ErrNotFound = errors.New("task_not_found")

func NewMemStore() *MemStore {
	return &MemStore{
		tasks: make([]*entities.Task, 0),
	}
}

func (s *MemStore) List() []*entities.Task {
	s.mx.Lock()
	defer s.mx.Unlock()

	return s.tasks[:]
}

func (s *MemStore) Create(t *entities.Task) {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.tasks = append(s.tasks, t)
}

func (s *MemStore) findIndexById(id int64) *int {
	var index *int

	for i, t := range s.tasks {
		if t.ID == id {
			index = &i
			break
		}
	}

	return index
}

func (s *MemStore) Update(t *entities.Task) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	if i := s.findIndexById(t.ID); i != nil {
		s.tasks[*i] = t
		return nil
	}

	return ErrNotFound
}

func (s *MemStore) Delete(id int64) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	if id < 0 || id >= int64(len(s.tasks)) {
		return ErrNotFound
	}

	i := s.findIndexById(id)
	if i != nil {
		s.tasks = append(s.tasks[:*i], s.tasks[*i+1:]...)
		return nil
	}
	return ErrNotFound
}

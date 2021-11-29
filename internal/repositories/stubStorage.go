package todorepo

import (
	"errors"

	"github.com/amr9mohamed/todoAPI/internal/core/domain"
)

type stubStorage struct {
	todos []domain.Todo
}

func NewStubStorage() *stubStorage {
	return &stubStorage{todos: []domain.Todo{}}
}

func (s *stubStorage) Get(id uint64) (domain.Todo, error) {
	for _, t := range s.todos {
		if t.ID == id {
			return t, nil
		}
	}
	return domain.Todo{}, errors.New("todo not found")
}

func (s *stubStorage) List() ([]domain.Todo,error) {
	return s.todos, nil
}

func (s *stubStorage) Delete(id uint64) error {
	for i, t := range s.todos {
		if t.ID == id {
			s.todos = append(s.todos[:i], s.todos[i+1:]...)
			return nil
		}
	}
	return errors.New("todo not found to be deleted")
}

func (s *stubStorage) Add(newTodo domain.Todo) error {
	s.todos = append(s.todos, newTodo)
	return nil
}

func (s *stubStorage) Edit(id uint64, editedTodo domain.Todo) error {
	for i, t := range s.todos {
		if t.ID == id {
			s.todos[i] = editedTodo
			return nil
		}
	}
	return errors.New("todo not found to be edited")
}

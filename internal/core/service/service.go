package todosrv

import (
	"errors"
	"github.com/amr9mohamed/todoAPI/internal/core/domain"
	"github.com/amr9mohamed/todoAPI/internal/core/ports"
)

type service struct {
	todoRepository ports.TodoRepository
}

func New(todoRepository ports.TodoRepository) *service {
	return &service{
		todoRepository: todoRepository,
	}
}

func (s *service) Get(id uint64) (domain.Todo, error) {
	todo, err := s.todoRepository.Get(id)
	if err != nil {
		return domain.Todo{}, errors.New("todo not found")
	}
	return todo, nil
}

func (s *service) List() []domain.Todo {
	return s.todoRepository.List()
}

func (s *service) Delete(id uint64) error {
	err := s.todoRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Add(newTodo domain.Todo) error {
	err := s.todoRepository.Add(newTodo)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Edit(id uint64, editedTodo domain.Todo) error {
	err := s.todoRepository.Edit(id, editedTodo)
	if err != nil {
		return err
	}
	return nil
}

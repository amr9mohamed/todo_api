package todosrv

import (
	"errors"
	"strconv"

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

func (s *service) Get(id string) (domain.Todo, error) {
	id_int, _ := strconv.ParseInt(id, 10, 64)
	todo, err := s.todoRepository.Get(uint64(id_int))
	if err != nil {
		return domain.Todo{}, errors.New("todo not found")
	}
	return todo, nil
}

func (s *service) List() []domain.Todo {
	return s.todoRepository.List()
}

func (s *service) Delete(id string) error {
	// todo fix conversion later
	id_int, _ := strconv.ParseInt(id, 10, 64)
	err := s.todoRepository.Delete(uint64(id_int))
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

func (s *service) Edit(id string, editedTodo domain.Todo) error {
	id_int, _ := strconv.ParseInt(id, 10, 64)
	err := s.todoRepository.Edit(uint64(id_int), editedTodo)
	if err != nil {
		return err
	}
	return nil
}

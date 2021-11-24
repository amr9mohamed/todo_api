package ports

import "github.com/amr9mohamed/todoAPI/internal/core/domain"

type TodoService interface {
	Get(id string) (domain.Todo, error)
	List() []domain.Todo
	Delete(id string) error
	Add(domain.Todo) error
	Edit(id string, editedTodo domain.Todo) error
}

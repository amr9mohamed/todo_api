package ports

import "github.com/amr9mohamed/todoAPI/internal/core/domain"

type TodoService interface {
	Get(id uint64) (domain.Todo, error)
	List() []domain.Todo
	Delete(id uint64) error
	Add(domain.Todo) error
	Edit(id uint64, editedTodo domain.Todo) error
}

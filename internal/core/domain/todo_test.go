package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTodo(t *testing.T) {
	t.Run("test create todo", func(t *testing.T) {
		todo := NewTodo(1, "Hello world", true)
		assert.Equal(t, uint64(1), todo.ID)
		assert.Equal(t, "Hello world", todo.Description)
		assert.Equal(t, true, todo.Completed)
	})
}

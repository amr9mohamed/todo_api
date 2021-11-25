package todosrv

import (
	"errors"
	"testing"

	"github.com/amr9mohamed/todoAPI/internal/core/domain"
	"github.com/amr9mohamed/todoAPI/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	mockRepo, service := setUp()

	type want struct {
		todo domain.Todo
		err  error
	}

	type args struct {
		id string
	}

	todo := domain.Todo{
		ID:          1,
		Description: "Hello",
		Completed:   true,
	}

	tests := []struct {
		name     string
		args     args
		want     want
		mockFunc func()
	}{
		{
			name:     "should get successfully",
			args:     args{id: "1"},
			want:     want{todo: todo, err: nil},
			mockFunc: func() { mockRepo.On("Get", uint64(1)).Return(todo, nil) },
		},
		{
			name:     "should return empty todo and err",
			args:     args{id: "2"},
			want:     want{todo: domain.Todo{}, err: errors.New("todo not found")},
			mockFunc: func() { mockRepo.On("Get", uint64(2)).Return(domain.Todo{}, errors.New("todo not found")) },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			todo, err := service.Get(tt.args.id)
			expected := tt.want
			assert.Equal(t, expected.todo, todo)
			assert.Equal(t, expected.err, err)

		})
	}
}

func TestList(t *testing.T) {
	mockRepo, service := setUp()

	type want struct {
		todos []domain.Todo
	}

	todos := []domain.Todo{
		{
			ID:          1,
			Description: "Hello",
			Completed:   false,
		},
		{
			ID:          2,
			Description: "world",
			Completed:   true,
		},
	}

	tests := []struct {
		name     string
		want     want
		mockFunc func()
	}{
		{
			name: "list all todo",
			want: want{todos: todos},
			mockFunc: func() {
				mockRepo.On("List").Return(todos).Once()
			},
		},
		{
			name: "return empty list",
			want: want{todos: []domain.Todo{}},
			mockFunc: func() {
				mockRepo.On("List").Return([]domain.Todo{}).Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			todos := service.List()
			expected := tt.want.todos
			assert.Equal(t, expected, todos)
		})
	}
}

func TestDelete(t *testing.T) {
	mockRepo, service := setUp()

	type args struct {
		id string
	}

	type want struct {
		err error
	}

	tests := []struct {
		name     string
		args     args
		want     want
		mockFunc func()
	}{
		{
			name: "found and deleted",
			args: args{id: "1"},
			want: want{err: nil},
			mockFunc: func() {
				mockRepo.On("Delete", uint64(1)).Return(nil).Once()
			},
		},
		{
			name: "not found and should return error",
			args: args{id: "1"},
			want: want{err: errors.New("error")},
			mockFunc: func() {
				mockRepo.On("Delete", uint64(1)).Return(errors.New("error")).Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			err := service.Delete(tt.args.id)
			expected := tt.want.err
			assert.Equal(t, expected, err)
		})
	}
}

func TestAdd(t *testing.T) {
	mockRepo, service := setUp()

	type args struct {
		todo domain.Todo
	}

	type want struct {
		err error
	}

	todo := domain.Todo{
		ID:          1,
		Description: "Hello world",
		Completed:   true,
	}

	tests := []struct {
		name     string
		args     args
		want     want
		mockFunc func()
	}{
		{
			name: "added successfully",
			args: args{todo: todo},
			want: want{err: nil},
			mockFunc: func() {
				mockRepo.On("Add", todo).Return(nil).Once()
			},
		},
		{
			name: "failed to add",
			args: args{todo: todo},
			want: want{err: errors.New("error")},
			mockFunc: func() {
				mockRepo.On("Add", todo).Return(errors.New("error")).Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			err := service.Add(tt.args.todo)
			expected := tt.want.err
			assert.Equal(t, expected, err)
		})
	}
}

func TestEdit(t *testing.T) {
	mockRepo, service := setUp()

	type args struct {
		id string
		todo domain.Todo
	}

	type want struct {
		err error
	}

	todo := domain.Todo{
		ID: 1,
		Description: "Hello world",
		Completed: true,
	}

	tests := []struct{
		name string
		args args
		want want
		mockFunc func()
	}{
		{
			name: "return nil",
			args: args{id: "1", todo: todo},
			want: want{err: nil},
			mockFunc: func() {
				mockRepo.On("Edit", uint64(1), todo).Return(nil).Once()
			},
		},
		{
			name: "return error",
			args: args{id: "1", todo: todo},
			want: want{err: errors.New("error")},
			mockFunc: func() {
				mockRepo.On("Edit", uint64(1), todo).Return(errors.New("error")).Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			err := service.Edit(tt.args.id, tt.args.todo)
			expected := tt.want.err
			assert.Equal(t, expected, err)
		})
	}
}

func setUp() (*mocks.TodoRepository, *service) {
	mockRepo := new(mocks.TodoRepository)
	service := New(mockRepo)
	return mockRepo, service
}

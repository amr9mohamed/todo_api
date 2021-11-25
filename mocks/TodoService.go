// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	domain "github.com/amr9mohamed/todoAPI/internal/core/domain"
	mock "github.com/stretchr/testify/mock"
)

// TodoService is an autogenerated mock type for the TodoService type
type TodoService struct {
	mock.Mock
}

// Add provides a mock function with given fields: _a0
func (_m *TodoService) Add(_a0 domain.Todo) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Todo) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *TodoService) Delete(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Edit provides a mock function with given fields: id, editedTodo
func (_m *TodoService) Edit(id string, editedTodo domain.Todo) error {
	ret := _m.Called(id, editedTodo)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, domain.Todo) error); ok {
		r0 = rf(id, editedTodo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: id
func (_m *TodoService) Get(id string) (domain.Todo, error) {
	ret := _m.Called(id)

	var r0 domain.Todo
	if rf, ok := ret.Get(0).(func(string) domain.Todo); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.Todo)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields:
func (_m *TodoService) List() []domain.Todo {
	ret := _m.Called()

	var r0 []domain.Todo
	if rf, ok := ret.Get(0).(func() []domain.Todo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Todo)
		}
	}

	return r0
}

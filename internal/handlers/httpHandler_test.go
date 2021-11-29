package todohdl

import (
	"encoding/json"
	"errors"
	"github.com/amr9mohamed/todoAPI/internal/core/domain"
	"github.com/amr9mohamed/todoAPI/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestList(t *testing.T) {
	mockService := new(mocks.TodoService)
	handler := NewHTTPHandler(mockService)
	router := gin.Default()
	router.GET("/todo", handler.List)

	t.Run("test list api with actual data", func(t *testing.T) {
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
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/todo", nil)
		mockFunc := func() { mockService.On("List").Return(todos, nil).Once() }
		mockFunc()
		router.ServeHTTP(w, req)
		var resTodos []domain.Todo
		if err := json.Unmarshal(w.Body.Bytes(), &resTodos); err != nil {
			return
		}
		assert.Equal(t, 200, w.Code)
		assert.Equal(t, todos, resTodos)
	})
	t.Run("return error", func(t *testing.T) {
		var todos  []domain.Todo
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/todo", nil)
		mockFunc := func() { mockService.On("List").Return(todos, errors.New("error")).Once() }
		mockFunc()
		router.ServeHTTP(w, req)
		var resTodos []domain.Todo
		if err := json.Unmarshal(w.Body.Bytes(), &resTodos); err != nil {
			return
		}
		assert.Equal(t, 404, w.Code)
		assert.Equal(t, todos, resTodos)
	})
}

func TestGet(t *testing.T) {
	mockService := new(mocks.TodoService)
	handler := NewHTTPHandler(mockService)
	router := gin.Default()
	router.GET("/todo/:id", handler.Get)

	t.Run("test get api", func(t *testing.T) {
		todo := domain.NewTodo(1, "Hello world", true)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/todo/1", nil)
		mockFunc := func() { mockService.On("Get", uint64(1)).Return(todo, nil).Once() }
		mockFunc()
		router.ServeHTTP(w, req)
		var resTodo domain.Todo
		if err := json.Unmarshal(w.Body.Bytes(), &resTodo); err != nil {
			return
		}
		assert.Equal(t, 200, w.Code)
		assert.Equal(t, todo, resTodo)
	})
	t.Run("should return error", func(t *testing.T) {
		todo := domain.Todo{}
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/todo/1", nil)
		mockFunc := func() { mockService.On("Get", uint64(1)).Return(todo, errors.New("error")).Once() }
		mockFunc()
		router.ServeHTTP(w, req)
		var resTodo domain.Todo
		if err := json.Unmarshal(w.Body.Bytes(), &resTodo); err != nil {
			return
		}
		assert.Equal(t, 404, w.Code)
		assert.Equal(t, todo, resTodo)
	})
}

func TestDelete(t *testing.T) {
	mockService := new(mocks.TodoService)
	handler := NewHTTPHandler(mockService)
	router := gin.Default()
	router.DELETE("/todo/:id", handler.Delete)

	t.Run("Should delete ", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/todo/1", nil)
		mockFunc := func() { mockService.On("Delete", uint64(1)).Return(nil).Once() }
		mockFunc()
		router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
	})
	t.Run("should return error", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/todo/1", nil)
		mockFunc := func() { mockService.On("Delete", uint64(1)).Return(errors.New("err")).Once() }
		mockFunc()
		router.ServeHTTP(w, req)
		assert.Equal(t, 404, w.Code)
	})
}

func TestAdd(t *testing.T) {
	mockService := new(mocks.TodoService)
	handler := NewHTTPHandler(mockService)
	router := gin.Default()
	router.POST("/todo", handler.Add)

	t.Run("Should add", func(t *testing.T) {
		todo := domain.NewTodo(1, "hello world", true)
		bytesTodo, _ := json.Marshal(todo)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/todo", strings.NewReader(string(bytesTodo)))
		mockFunc := func() { mockService.On("Add", todo).Return(nil).Once() }
		mockFunc()
		router.ServeHTTP(w, req)
		assert.Equal(t, 204, w.Code)
	})
	t.Run("should return err", func(t *testing.T) {
		todo := domain.NewTodo(1, "hello world", true)
		bytesTodo, _ := json.Marshal(todo)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/todo", strings.NewReader(string(bytesTodo)))
		mockFunc := func() { mockService.On("Add", todo).Return(errors.New("error")).Once() }
		mockFunc()
		router.ServeHTTP(w, req)
		assert.Equal(t, 409, w.Code)
	})
	t.Run("not well structured", func(t *testing.T) {
		todo := domain.NewTodo(1, "hello", false)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/todo", nil)
		mockFunc := func() { mockService.On("Add", todo).Return(errors.New("error")).Once() }
		mockFunc()
		router.ServeHTTP(w, req)
		assert.Equal(t, 422, w.Code)
	})
}

func TestEdit(t *testing.T) {
	mockService := new(mocks.TodoService)
	handler := NewHTTPHandler(mockService)
	router := gin.Default()
	router.PATCH("/todo/:id", handler.Edit)

	t.Run("Should return error could not edit", func(t *testing.T) {
		todo := domain.NewTodo(1, "hello world", true)
		bytesTodo, _ := json.Marshal(todo)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PATCH", "/todo/1", strings.NewReader(string(bytesTodo)))
		mockFunc := func() {
			mockService.On("Get", uint64(1)).Return(todo, nil).Once()
			mockService.On("Edit", uint64(1), todo).Return(errors.New("error")).Once()
		}
		mockFunc()
		router.ServeHTTP(w, req)
		assert.Equal(t, 400, w.Code)
	})
	t.Run("should return err not found", func(t *testing.T) {
		todo := domain.NewTodo(1, "hello world", true)
		bytesTodo, _ := json.Marshal(todo)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PATCH", "/todo/1", strings.NewReader(string(bytesTodo)))
		mockFunc := func() {
			mockService.On("Get", uint64(1)).Return(todo, errors.New("error")).Once()
			mockService.On("Edit", uint64(1), todo).Return(nil).Once()
		}
		mockFunc()
		router.ServeHTTP(w, req)
		assert.Equal(t, 404, w.Code)
	})
	t.Run("not well structured", func(t *testing.T) {
		todo := domain.NewTodo(1, "hello world", true)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PATCH", "/todo/1", nil)
		mockFunc := func() {
			mockService.On("Get", uint64(1)).Return(todo, nil).Once()
			mockService.On("Edit", uint64(1), todo).Return(nil).Once()
		}
		mockFunc()
		router.ServeHTTP(w, req)
		assert.Equal(t, 422, w.Code)
	})
	t.Run("Should edit", func(t *testing.T) {
		todo := domain.NewTodo(1, "hello world", true)
		bytesTodo, _ := json.Marshal(todo)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PATCH", "/todo/1", strings.NewReader(string(bytesTodo)))
		mockFunc := func() {
			mockService.On("Get", uint64(1)).Return(todo, nil).Once()
			mockService.On("Edit", uint64(1), todo).Return(nil).Once()
		}
		mockFunc()
		router.ServeHTTP(w, req)
		assert.Equal(t, 204, w.Code)
	})
}

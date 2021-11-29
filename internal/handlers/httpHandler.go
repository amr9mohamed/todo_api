package todohdl

import (
	"net/http"

	"github.com/amr9mohamed/todoAPI/internal/core/domain"
	"github.com/amr9mohamed/todoAPI/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	todoService ports.TodoService
}

func NewHTTPHandler(todoService ports.TodoService) *HTTPHandler {
	return &HTTPHandler{
		todoService: todoService,
	}
}

func (h *HTTPHandler) Get(c *gin.Context) {
	var todo domain.Todo
	if err := c.ShouldBindUri(&todo); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, err)
		return
	}
	todo, err := h.todoService.Get(todo.ID)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, todo)
}

func (h *HTTPHandler) List(c *gin.Context) {
	todos := h.todoService.List()
	c.IndentedJSON(http.StatusOK, todos)
}

func (h *HTTPHandler) Delete(c *gin.Context) {
	var todo domain.Todo
	if err := c.ShouldBindUri(&todo); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, err)
		return
	}
	if err := h.todoService.Delete(todo.ID); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, nil)
}

func (h *HTTPHandler) Add(c *gin.Context) {
	body := domain.Todo{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, err)
		return
	}
	if err := h.todoService.Add(body); err != nil {
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	c.IndentedJSON(http.StatusNoContent, nil)
}

func (h *HTTPHandler) Edit(c *gin.Context) {
	var todo domain.Todo
	if err := c.ShouldBindUri(&todo); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, err)
		return
	}
	_, err := h.todoService.Get(todo.ID)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
		return
	}
	body := domain.Todo{}
	body.ID = todo.ID
	if err := c.ShouldBindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, err)
		return
	}
	if err := h.todoService.Edit(todo.ID, body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusNoContent, nil)
}

package controllers

import (
	"net/http"

	"github.com/amr9mohamed/todoAPI/models"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetTodos(c *gin.Context) {
	todo := models.Todo{}
	todos, err := todo.GetTodos(s.DB)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusOK, todos)
}

func (s *Server) PostTodo(c *gin.Context) {
	var newTodo models.Todo
	if err := c.BindJSON(&newTodo); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, err)
	}
	todo := models.Todo{}
	err := todo.PostTodo(s.DB, newTodo)
	if err != nil {
		c.IndentedJSON(http.StatusConflict, err)
	}
	c.IndentedJSON(http.StatusOK, nil)
}

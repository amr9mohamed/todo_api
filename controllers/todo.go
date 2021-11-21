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
	newTodo := models.Todo{}
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

func (s *Server) DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	todo := models.Todo{}
	if err := todo.DeleteTodo(s.DB, id); err == nil {
		c.IndentedJSON(http.StatusNoContent, nil)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found to be deleted"})
}

func (s *Server) GetTodo(c *gin.Context) {
	id := c.Param("id")
	todo := models.Todo{}
	targetTodo, err := todo.GetTodo(s.DB, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, nil)
		return
	}
	c.IndentedJSON(http.StatusOK, targetTodo)
}

func (s *Server) EditTodo(c *gin.Context) {
	id := c.Param("id")
	editedTodo := models.Todo{}
	if err := c.ShouldBindJSON(&editedTodo); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, err)
		return
	}
	todo := models.Todo{}
	if err := todo.EditTodo(s.DB, id, editedTodo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusNoContent, nil)
}

package server

import (
	todosrv "github.com/amr9mohamed/todoAPI/internal/core/service"
	todohdl "github.com/amr9mohamed/todoAPI/internal/handlers"
	todorepo "github.com/amr9mohamed/todoAPI/internal/repositories"
	"github.com/gin-gonic/gin"
)

func RunHTTP() {
	todoRepo := todorepo.NewPsqlDB("postgres", "amrmohamed", "password", "5432", "todo-postgres", "users")
	// todoRepo := todorepo.NewStubStorage()
	todoSrv := todosrv.New(todoRepo)
	todoHandler := todohdl.NewHTTPHandler(todoSrv)

	router := gin.Default()
	todoRouter := router.Group("/todo")
	{
		todoRouter.GET("", todoHandler.List)
		todoRouter.POST("", todoHandler.Add)
		todoRouter.GET("/:id", todoHandler.Get)
		todoRouter.PATCH("/:id", todoHandler.Edit)
		todoRouter.DELETE("/:id", todoHandler.Delete)
	}

	router.Run(":8080")
}

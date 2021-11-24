package server

import (
	"log"

	todosrv "github.com/amr9mohamed/todoAPI/internal/core/service"
	todohdl "github.com/amr9mohamed/todoAPI/internal/handlers"
	todorepo "github.com/amr9mohamed/todoAPI/internal/repositories"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func RunHTTP() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Print("error loading .env file")
		panic(err)
	}
	todoRepo := todorepo.NewPsqlDB(viper.GetString("DB_DRIVER"), viper.GetString("DB_USER"), viper.GetString("DB_PASSWORD"), viper.GetString("DB_PORT"), viper.GetString("DB_HOST"), viper.GetString("DB_NAME"))
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

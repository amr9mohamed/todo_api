package controllers

func (s *Server) initializeRoutes() {
	//Todo routes
	todoRouter := s.Router.Group("/todo")
	{
		todoRouter.GET("", s.GetTodos)
		todoRouter.POST("", s.PostTodo)
		todoRouter.GET("/:id", s.GetTodo)
		todoRouter.PATCH("/:id", s.EditTodo)
		todoRouter.DELETE("/:id", s.DeleteTodo)
	}
}

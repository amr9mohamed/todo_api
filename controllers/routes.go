package controllers

func (s *Server) initializeRoutes() {
	//Todo routes
	s.Router.GET("/todo", s.GetTodos)
	s.Router.POST("/todo", s.PostTodo)
}

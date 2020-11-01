package server

import "net/http"

func (s *Server) routes() {
	// serve static files
	s.router.StaticFS("/public/", http.Dir("./public"))
	s.router.StaticFile("/favicon.ico", "./public/favicon-32x32.png")
	s.router.StaticFile("/", "./public/index.html")

	///////////
	// USERS //
	///////////
	s.router.GET("/users/:userID", s.users.GetUser())
	s.router.PUT("/users/:userID", s.users.UpdateUser())
	s.router.DELETE("/users/:userID", s.users.DeleteUser())
	s.router.POST("/users/", s.users.CreateUser())
}

package server

import (
	"fmt"
	"log"
	"os"

	"github.com/OJOMB/mvsc-microservice/config"
	"github.com/OJOMB/mvsc-microservice/controllers"
	"github.com/gin-gonic/gin"
)

// Server encapsulates the shared deps
type Server struct {
	router  *gin.Engine
	users   controllers.UsersInterface
	threads controllers.ThreadsInterface
	posts   controllers.PostsInterface
	topics  controllers.TopicsInterface
	config  *config.Config
}

// New returns a new Server instance
func New(router *gin.Engine, config *config.Config) (s *Server) {
	s = &Server{
		router:  router,
		users:   controllers.Users,
		threads: controllers.Threads,
		posts:   controllers.Posts,
		topics:  controllers.Topics,
		config:  config,
	}

	// setup middlewares
	s.middlewares()
	// setup all routes
	s.routes()

	return
}

// ListenAndServe For Listening and Serving
func (s *Server) ListenAndServe() {
	addr := fmt.Sprintf("%s:%d", s.config.IP, s.config.Port)
	log.Fatal(s.router.Run(addr))
}

// StartApp takes some config and handles the initial startup process
func StartApp(env string) {
	// get the config
	cnfg := config.ConfigMap[env]

	// get the logger
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Printf("Server is starting...")

	// get the router
	router := gin.Default()

	// Instantiate server with shared dependencies
	s := New(router, &cnfg)

	s.ListenAndServe()
}

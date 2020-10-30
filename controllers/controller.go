package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/OJOMB/mvsc-microservice/config"
	"github.com/gorilla/mux"
)

// Controller encapsulates the shared deps
type Controller struct {
	router *mux.Router
	logger *log.Logger
	config *config.Config
}

// New returns a new Server instance
func New(router *mux.Router, logger *log.Logger, config *config.Config) *Controller {
	router.Use(contentTypeMiddleware)
	var c *Controller = &Controller{
		config: config,
		router: router,
		logger: logger,
	}
	c.routes()
	return c
}

func contentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		},
	)
}

func (c *Controller) routes() {
	fs := http.FileServer(http.Dir("./public"))
	handlePublic := http.StripPrefix("/public/", fs)
	c.router.PathPrefix("/public/").Handler(handlePublic)

	c.router.HandleFunc("/", c.handleIndex())
	c.router.HandleFunc("/users/{id}", c.GetUser()).Methods("GET")
}

func (c *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.router.ServeHTTP(w, r)
}

// ListenAndServe Listens and serves on the address specified in the controller config
func (c *Controller) ListenAndServe() {
	addr := fmt.Sprintf("%s:%d", c.config.IP, c.config.Port)
	c.logger.Printf("Server listening on: %s", addr)
	c.logger.Fatal(http.ListenAndServe(addr, c))
}

func (c *Controller) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/index.html")
		c.logger.Print("/ served")
	}
}

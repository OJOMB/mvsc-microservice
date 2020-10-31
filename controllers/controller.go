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
	var c *Controller = &Controller{
		config: config,
		router: router,
		logger: logger,
	}

	// setup middlewares
	router.Use(c.contentTypeMiddleware)
	router.Use(c.logURLServedMiddleWare)

	// setup all routes
	c.routes()

	return c
}

func (c *Controller) contentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/" {
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				c.logger.Printf("Content-Type: %s", w.Header().Get("Content-Type"))
			} else {
				c.logger.Print("set the content-type to json")
				w.Header().Set("Content-Type", "application/json")
			}
			next.ServeHTTP(w, r)
		},
	)
}

func (c *Controller) logURLServedMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
			c.logger.Printf("%s served", r.URL.Path)
		},
	)
}

func (c *Controller) routes() {
	fs := http.FileServer(http.Dir("./public"))
	handlePublic := http.StripPrefix("/public/", fs)
	c.router.PathPrefix("/public/").Handler(handlePublic)

	c.router.HandleFunc("/", c.handleIndex())
	c.router.HandleFunc("/users/{id}", c.GetUser()).Methods("GET")
	c.router.HandleFunc("/users", c.CreateUser()).Methods("POST")
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
	}
}

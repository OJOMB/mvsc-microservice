package main

import (
	"flag"
	"log"
	"os"

	"github.com/OJOMB/mvsc-microservice/config"
	"github.com/OJOMB/mvsc-microservice/controllers"
	"github.com/gorilla/mux"
)

var env = flag.String("env", "dev", "The environment in which the server is running ['dev', 'test', 'production']")

func main() {
	flag.Parse()

	// get the logger
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Printf("Server is starting...")

	// get the config
	cnfg := config.ConfigMap[*env]

	//get the router
	r := mux.NewRouter()

	// Instantiate server with shared dependencies
	c := controllers.New(r, logger, &cnfg)

	c.ListenAndServe()
}

package main

import (
	"flag"

	"github.com/OJOMB/mvsc-microservice/server"
)

var env = flag.String("env", "dev", "The environment in which the server is running ['dev', 'test', 'production']")

func main() {
	flag.Parse()
	server.StartApp(*env)
}

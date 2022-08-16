package main

import (
	"log"
	"os"

	"github.com/terrylin13/go-api-docker/pkg/router"
)

func main() {
	e := router.InitRouter()

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	log.Fatal(e.Run(":" + httpPort))
}

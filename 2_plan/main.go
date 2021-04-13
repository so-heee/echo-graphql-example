package main

import (
	"log"
	"net/http"
	"os"

	"github.com/so-heee/graphql-example/plan02/server"
)

const defaultPort = "8080"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	c := server.Config{Logging: true}
	s := server.NewServer(c)
	r := s.Router()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

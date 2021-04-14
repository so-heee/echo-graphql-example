package main

import (
	"log"
	"net/http"
	"os"

	"github.com/so-heee/graphql-example/2_plan/infrastructure"
)

const defaultPort = "8080"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	c := infrastructure.Config{Logging: true}
	s := infrastructure.NewServer(c)
	r := s.Router()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

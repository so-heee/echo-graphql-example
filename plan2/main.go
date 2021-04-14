package main

import (
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/so-heee/graphql-example/plan2/infrastructure"
)

const defaultPort = "8080"
const defaultDsn = "user:password@tcp(mysql):3306)/sample?charset=utf8mb4&parseTime=True&loc=Local"

type Env struct {
	MysqlUser     string `required:"true" split_words:"true"`
	MysqlPassword string `required:"true" split_words:"true"`
	MysqlDatabase string `required:"true" split_words:"true"`
	MysqlHost     string `required:"true" split_words:"true"`
}

func main() {

	var env Env
	if err := envconfig.Process("", &env); err != nil {
		log.Fatalf("[ERROR] Failed to process env: %s", err.Error())
	}

	dsn := env.MysqlUser + ":" + env.MysqlPassword + "@tcp(" + env.MysqlHost + ":3306)/" + env.MysqlDatabase + "?charset=utf8mb4&parseTime=True&loc=Local"

	c := infrastructure.Config{Logging: true}
	s := infrastructure.NewServer(c)
	r := s.Router(dsn)

	log.Printf("connect to http://localhost:8080/prayground for GraphQL playground")
	log.Fatal(http.ListenAndServe(":8080", r))
}

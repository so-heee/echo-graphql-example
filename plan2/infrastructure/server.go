package infrastructure

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kelseyhightower/envconfig"
)

// Server represents main server structure
type Server struct {
	router *chi.Mux
	config Config
}

type Config struct {
	Logging bool
}

// NewServer returns server with initialized router
func NewServer(c Config) *Server {
	return &Server{
		router: chi.NewRouter(),
		config: c,
	}
}

const defaultPort = "8080"
const defaultDsn = "user:password@tcp(mysql):3306)/sample?charset=utf8mb4&parseTime=True&loc=Local"

type Env struct {
	MysqlUser     string `required:"true" split_words:"true"`
	MysqlPassword string `required:"true" split_words:"true"`
	MysqlDatabase string `required:"true" split_words:"true"`
	MysqlHost     string `required:"true" split_words:"true"`
}

var env Env

func init() {
}

func Run() {
	if err := envconfig.Process("", &env); err != nil {
		log.Fatalf("[ERROR] Failed to process env: %s", err.Error())
	}

	dsn := env.MysqlUser + ":" + env.MysqlPassword + "@tcp(" + env.MysqlHost + ":3306)/" + env.MysqlDatabase + "?charset=utf8mb4&parseTime=True&loc=Local"

	c := Config{Logging: true}
	s := NewServer(c)
	r := s.Router(dsn)

	log.Printf("connect to http://localhost:8080/prayground for GraphQL playground")
	log.Fatal(http.ListenAndServe(":8080", r))
}

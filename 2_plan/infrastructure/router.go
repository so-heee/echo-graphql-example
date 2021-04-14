package infrastructure

import (
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
	"github.com/so-heee/graphql-example/2_plan/graph/generated"
	"github.com/so-heee/graphql-example/2_plan/graph/resolver"

	"github.com/so-heee/graphql-example/2_plan/interfaces/database"
)

func (s *Server) Router() *chi.Mux {

	// Use JSON logger
	customLogger := logrus.New()
	customLogger.Formatter = &logrus.JSONFormatter{
		// disable, as we set our own
		DisableTimestamp: true,
	}

	dsn := "user:password@tcp(localhost:3306)/sample?charset=utf8mb4&parseTime=True&loc=Local"
	repo, err := database.NewRepository(dsn)
	if err != nil {
		fmt.Errorf("failed to create repository: %+v", err)
	}
	resolver := resolver.NewResolver(repo)

	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)

	s.router.Use(render.SetContentType(render.ContentTypeJSON))

	graphqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	playgroundHandler := playground.Handler("GraphQL Playground", "/query")

	s.router.Handle("/query", graphqlHandler)
	s.router.Handle("/playground", playgroundHandler)

	return s.router
}

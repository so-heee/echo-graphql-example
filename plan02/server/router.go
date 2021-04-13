package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
	"github.com/so-heee/graphql-example/plan02/graph/generated"
	"github.com/so-heee/graphql-example/plan02/graph/resolver"
)

func (s *Server) Router() *chi.Mux {

	// Use JSON logger
	customLogger := logrus.New()
	customLogger.Formatter = &logrus.JSONFormatter{
		// disable, as we set our own
		DisableTimestamp: true,
	}

	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)

	s.router.Use(render.SetContentType(render.ContentTypeJSON))

	graphqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))
	playgroundHandler := playground.Handler("GraphQL Playground", "/query")

	s.router.Handle("/query", graphqlHandler)
	s.router.Handle("/playground", playgroundHandler)

	return s.router
}

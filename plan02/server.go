package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/so-heee/graphql-example/plan02/graph/generated"
	"github.com/so-heee/graphql-example/plan02/graph/resolver"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	graphqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))
	playgroundHandler := playground.Handler("GraphQL Playground", "/query")

	e.POST("/query", echo.WrapHandler(graphqlHandler))

	e.GET("/playground", echo.WrapHandler(playgroundHandler))

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

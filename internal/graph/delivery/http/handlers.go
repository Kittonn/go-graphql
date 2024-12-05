package http

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Kittonn/go-graphql/internal/graph"
	"github.com/Kittonn/go-graphql/internal/graph/generated"
	"github.com/labstack/echo/v4"
)

type graphqlHandlers struct{}

func NewGraphQLHandlers() graph.Handlers {
	return &graphqlHandlers{}
}

func (g *graphqlHandlers) PlaygroundHandler() echo.HandlerFunc {
	graphQLPlayground := playground.Handler("GraphQL", "/graphql")

	return func(c echo.Context) error {
		graphQLPlayground.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}

func (g *graphqlHandlers) GraphQLHandler() echo.HandlerFunc {
	executableSchema := generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}})
	graphQLHandler := handler.NewDefaultServer(executableSchema)

	return func(c echo.Context) error {
		graphQLHandler.ServeHTTP(c.Response(), c.Request())

		return nil
	}
}

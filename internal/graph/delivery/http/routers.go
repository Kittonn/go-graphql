package http

import (
	"github.com/Kittonn/go-graphql/internal/graph"
	"github.com/labstack/echo/v4"
)

func MapGraphQLRoutes(graphqlsGroup *echo.Group, graphqlHandlers graph.Handlers) {
	graphqlsGroup.GET("/playground", graphqlHandlers.PlaygroundHandler())
	graphqlsGroup.POST("/graphql", graphqlHandlers.GraphQLHandler())
}

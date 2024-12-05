package app

import graphQLHttp "github.com/Kittonn/go-graphql/internal/graph/delivery/http"

func (a *app) MapHandlers() error {
	baseGroup := a.echo.Group("")

	graphqlHandlers := graphQLHttp.NewGraphQLHandlers()
	graphQLHttp.MapGraphQLRoutes(baseGroup, graphqlHandlers)

	return nil
}

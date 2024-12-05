package graph

import "github.com/labstack/echo/v4"

type Handlers interface {
	GraphQLHandler() echo.HandlerFunc
	PlaygroundHandler() echo.HandlerFunc
}

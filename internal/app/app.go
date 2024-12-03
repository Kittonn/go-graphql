package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Kittonn/go-graphql/config"
	"github.com/labstack/echo/v4"
)

type app struct {
	echo   *echo.Echo
	config *config.Config
}

func NewApp(config *config.Config) *app {
	return &app{
		echo:   echo.New(),
		config: config,
	}
}

func (a *app) Run() error {
	if err := a.MapHandlers(); err != nil {
		return err
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	port := fmt.Sprintf(":%d", a.config.Port)
	go func() {
		if err := a.echo.Start(port); err != nil && err != http.ErrServerClosed {
			a.echo.Logger.Fatal("shutting down the server")
		}
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := a.echo.Shutdown(ctx); err != nil {
		a.echo.Logger.Fatal(err)
	}

	return nil
}

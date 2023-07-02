package http

import (
	"context"
	"net/http"
	"time"

	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/app/resolver"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo         *echo.Echo
	userResolver *resolver.UserResolver
}

func NewServer(userResolver *resolver.UserResolver) *Server {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return &Server{
		echo:         e,
		userResolver: userResolver,
	}
}

func (s *Server) Start(addr string) error {
	s.echo.GET("/apis/playground", PlaygroundHandler("/graphql"))
	s.echo.POST("/graphql", echo.WrapHandler(NewHandler(s.userResolver)))

	server := &http.Server{
		Addr:         addr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return s.echo.StartServer(server)
}

func (s *Server) Stop(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return s.echo.Shutdown(ctx)
}

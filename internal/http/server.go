package http

import (
	"context"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/handler"
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/app/resolver"
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/core/graph/generated"
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

	e.GET("/apis/playground", playgroundHandler())
	e.POST("/apis/graphql", echo.WrapHandler(graphqlHandler(userResolver)))

	return &Server{
		echo:         e,
		userResolver: userResolver,
	}
}

func (s *Server) Start(addr string) error {
	server := &http.Server{
		Addr:         addr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      s.echo,
	}

	return server.ListenAndServe()
}

func (s *Server) Stop(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return s.echo.Shutdown(ctx)
}

func graphqlHandler(userResolver *resolver.UserResolver) http.Handler {
	cfg := generated.Config{Resolvers: userResolver}
	h := handler.GraphQL(generated.NewExecutableSchema(cfg))

	return h
}

func playgroundHandler() echo.HandlerFunc {
	h := playground.Handler("GraphQL Playground", "/graphql")
	return func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}

package http

import (
	"context"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/core/graph"
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/core/graph/generated"
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/core/interfaces"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo         *echo.Echo
	rootResolver *graph.Resolver
}

func NewServer(userRepo interfaces.UserInterface, shopRepo interfaces.ShopInterface) *Server {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	rootResolver := graph.NewResolver(userRepo, shopRepo)

	e.GET("/apis/playground", playgroundHandler())
	e.POST("/apis/graphql", graphqlHandler(rootResolver))

	return &Server{
		echo:         e,
		rootResolver: rootResolver,
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

func graphqlHandler(resolver *graph.Resolver) echo.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	return func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}

func playgroundHandler() echo.HandlerFunc {
	h := playground.Handler("GraphQL Playground", "/apis/graphql")
	return func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}

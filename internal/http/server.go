package http

import (
	"context"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/app/resolver/root"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo         *echo.Echo
	rootResolver *root.RootResolver
}

func NewServer(rootResolver *root.RootResolver) *Server {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/apis/playground", playgroundHandler())
	// e.POST("/apis/graphql", echo.WrapHandler(graphqlHandler(rootResolver)))

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

// func graphqlHandler(userService *services.UserService, shopService *services.ShopService) http.Handler {
// srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
// h := handler.GraphQL(generated.NewExecutableSchema(cfg))
// return h
// }

func playgroundHandler() echo.HandlerFunc {
	h := playground.Handler("GraphQL Playground", "/")
	return func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}

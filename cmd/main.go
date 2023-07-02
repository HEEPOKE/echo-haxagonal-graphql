package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/sync/errgroup"
)

func main() {
	// Initialize MongoDB client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Initialize the repository
	db := client.Database("mydb")
	userRepo := &adapter.UserMongoRepository{DB: db}

	// Initialize the service
	userService := &domain.UserService{UserRepo: userRepo}

	// Initialize the Echo web framework
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize the GraphQL handler
	h := graph.NewHandler(userService)
	e.GET("/playground", graph.PlaygroundHandler("/graphql"))
	e.POST("/graphql", echo.WrapHandler(h))

	// Start the server
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g := new(errgroup.Group)
	g.Go(func() error {
		log.Println("Starting server at :8080")
		return e.StartServer(server)
	})

	// Graceful shutdown
	g.Go(func() error {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := e.Shutdown(ctx)
		if err != nil {
			log.Println("Server forced to shutdown:", err)
			return err
		}

		log.Println("Server shutdown gracefully")
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

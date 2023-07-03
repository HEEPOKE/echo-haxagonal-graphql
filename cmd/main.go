package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/app/resolver"
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/app/services"
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/domain/repositories"
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/http"
	"github.com/HEEPOKE/echo-haxagonal-graphql/pkg/config"
	"github.com/HEEPOKE/echo-haxagonal-graphql/pkg/database"
	"golang.org/x/sync/errgroup"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewMongoDB(cfg.MONGO_URL, cfg.DB_NAME)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := &repositories.UserRepository{DB: db.GetDatabase()}
	userResolver := &resolver.UserResolver{
		UserService: &services.UserService{
			UserRepo: userRepo,
		},
	}

	server := http.NewServer(userResolver)

	g := new(errgroup.Group)
	g.Go(func() error {
		log.Println("Starting server at :" + cfg.PORT)
		return server.Start(":" + cfg.PORT)
	})

	g.Go(func() error {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
		<-quit

		err := server.Stop(5 * time.Second)
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

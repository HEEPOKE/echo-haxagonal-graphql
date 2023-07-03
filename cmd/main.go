package main

import (
	"log"

	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/app/resolver"
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/app/services"
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/domain/repositories"
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/http"
	"github.com/HEEPOKE/echo-haxagonal-graphql/pkg/config"
	"github.com/HEEPOKE/echo-haxagonal-graphql/pkg/database"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := database.NewMongoDB(cfg.MONGO_URL, cfg.DB_NAME)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	userRepo := &repositories.UserRepository{DB: db.GetDatabase()}
	userService := services.NewUserService(userRepo)
	userResolver := resolver.NewUserResolver(userService)

	server := http.NewServer(userResolver)

	log.Printf("Starting server at :%s\n", cfg.PORT)
	err = server.Start(":" + cfg.PORT)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

package main

import (
	"context"
	"github.com/PhipattanachaiDev/golang_api-migration/configs"
	_ "github.com/PhipattanachaiDev/golang_api-migration/docs"
	"github.com/PhipattanachaiDev/golang_api-migration/internal/application"
	"github.com/PhipattanachaiDev/golang_api-migration/internal/infrastructure/db"
	"github.com/PhipattanachaiDev/golang_api-migration/internal/infrastructure/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.LoadConfig()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("userdb").Collection("users")
	repo := db.NewMongoUserRepository(collection)
	service := application.NewUserService(repo)

	// Start background goroutine
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("Stopping background task")
				return
			case <-time.After(10 * time.Second):
				users, _ := service.GetUsers()
				log.Printf("[Background Task] Total users: %d", len(users))
			}
		}
	}()

	r := http.SetupRouter(service)

	// Graceful shutdown
	go func() {
		if err := r.Run(":" + cfg.Port); err != nil {
			log.Fatalf("Server error: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down...")
	cancel()
	ctxShutdown, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()
	if err := client.Disconnect(ctxShutdown); err != nil {
		log.Printf("Mongo disconnect error: %v", err)
	}
	log.Println("Server exited cleanly")
}

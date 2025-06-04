package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI   string
	Port       string
	JWTSecret  string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	cfg := &Config{
		MongoURI:  os.Getenv("MONGO_URI"),
		Port:      os.Getenv("PORT"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}

	if cfg.MongoURI == "" || cfg.Port == "" || cfg.JWTSecret == "" {
		log.Fatal("Missing required environment variables")
	}

	return cfg
}
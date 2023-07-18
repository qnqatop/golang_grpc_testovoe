package config

import (
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"log"
)

var ApplicationConfig *Config

type Config struct {
	rdb *redis.Client
}

func NewConfig() *Config {
	_ = godotenv.Load()

	log.Print("initializing config")

	cfg := Config{}

	cfg.initRedis()

	return &cfg
}

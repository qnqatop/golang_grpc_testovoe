package config

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
)

func (cfg *Config) initRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	cfg.rdb = rdb

}

func (cfg *Config) GetRDb() *redis.Client {
	return cfg.rdb
}

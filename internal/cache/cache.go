package cache

import (
	"context"
	"time"

	"github.com/Campus-Hub/server/config"
	"github.com/Campus-Hub/server/pkg/logger"

	"github.com/go-redis/redis"
)

var (
	RedisClient *redis.Client
	Ctx         context.Context
)

// Setup the redis client.
// FINISHED TODO Setup Cache Module.
func Setup() {
	var ()

	RedisClient = redis.NewClient(&redis.Options{
		Addr:        config.Config.Redis.Host,
		Password:    config.Config.Redis.Password,
		DB:          0,
		IdleTimeout: time.Duration(config.Config.Redis.IdleTimeout),
	})

	Ctx = context.Background()

	// Test Redis Connection.
	_, err := RedisClient.Ping().Result()
	if err != nil {
		logger.Logger.Warn("Redis Setup Failed Error: ", err)
		return
	}
}

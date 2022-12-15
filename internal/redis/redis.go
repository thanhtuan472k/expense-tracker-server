package redis

import (
	"context"
	"encoding/json"
	"expense-tracker-server/internal/config"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"

	"github.com/logrusorgru/aurora"
)

var (
	store *redis.Client
)

// InitRedis ...
func InitRedis() {
	ctx := context.Background()
	cfg := config.GetENV().Redis

	store = redis.NewClient(&redis.Options{
		Addr:     cfg.URI,
		Password: cfg.Password,
		DB:       0, // use default DB
	})

	// Test
	_, err := store.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Cannot connect to redis", cfg.URI, err)
	}

	fmt.Println(aurora.Green("*** CONNECTED TO REDIS: " + cfg.URI))
}

// GetClient ...
func GetClient() *redis.Client {
	return store
}

// SetKeyValue ...
func SetKeyValue(key string, value interface{}, expiration time.Duration) error {
	ctx := context.Background()
	storeByte, _ := json.Marshal(value)
	r := store.Set(ctx, key, storeByte, expiration)
	return r.Err()
}

// GetValueByKey ...
func GetValueByKey(key string) string {
	ctx := context.Background()
	value, _ := store.Get(ctx, key).Result()
	return value
}

// GetJSON ...
func GetJSON(key string, result interface{}) (ok bool) {
	v := GetValueByKey(key)
	if v == "" {
		return false
	}
	if err := json.Unmarshal([]byte(v), result); err != nil {
		return false
	}
	return true
}

// DelKey ...
func DelKey(key string) error {
	ctx := context.Background()
	return store.Del(ctx, key).Err()
}

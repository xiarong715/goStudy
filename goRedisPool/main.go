package main

import (
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func main() {
	log.Printf("redis pool test")
	rdb = redis.NewClient(&redis.Options{
		Addr:     ":6379",
		Password: "",
		DB:       0,

		PoolSize:     15,
		MinIdleConns: 10,

		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolTimeout:  4 * time.Second,

		MaxRetries:      3,
		MinRetryBackoff: 8 * time.Millisecond,
		MaxRetryBackoff: 512 * time.Millisecond,

		MaxConnAge: 0 * time.Second,
	})
	defer rdb.Close()

}

// ReadRedis read redis
func ReadRedis(key string, value interface{}) {
	rdb.Set(rdb.Context(), key, value, 3*time.Second)
}

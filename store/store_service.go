package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// Define the struct wrapper around raw redis client
type StorageService struct {
	redisClient *redis.Client
}

//Top level declarations for the storeService and Redis context

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

//In a real world usage scenario, the cache duration should not have an expiration.
// An LRU policy config should be used where the values that are retrieved least often are evicted.
// The values should then be stored back in RDBMS when the cache is full and another value needs to be stored.

const CacheDuration = 6 * time.Hour

// Initialize the store service and return a store pointer
func InitStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}\n", pong)
	storeService.redisClient = redisClient
	return storeService
}

// We need to save the mapping between the original URL and the shortened URL.
func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	err := storeService.redisClient.Set(shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
}

// We need to retrieve the original URL from the shortened URL.
func RetrieveInitialUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to Retrieve InitialUrl | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return result
}

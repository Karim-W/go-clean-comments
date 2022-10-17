package database

import redis "github.com/go-redis/redis/v8"

var redisInstance *redis.Client

func GetRedisInstance() *redis.Client {
	mtx.Lock()
	defer mtx.Unlock()
	if redisInstance == nil {
		redisInstance = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	}
	return redisInstance
}

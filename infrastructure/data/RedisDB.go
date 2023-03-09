package data

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisDB struct {
	client *redis.Client
}

func NewRedisDB(options *redis.Options) (*RedisDB, error) {
	client := redis.NewClient(options)
	if err := client.Ping(context.TODO()).Err(); err != nil {
		return nil, err
	}
	return &RedisDB{client}, nil
}

func (redisDB *RedisDB) Set(key string, value string) error {
	return redisDB.client.Set(context.TODO(), key, value, 0).Err()
}

func (redisDB *RedisDB) Get(key string) (string, error) {
	return redisDB.client.Get(context.TODO(), key).Result()
}

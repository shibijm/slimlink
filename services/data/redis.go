package data

import (
	"context"
	"crypto/tls"
	"errors"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func InitRedisDb(options *redis.Options) {
	options.TLSConfig = &tls.Config{
		MinVersion: tls.VersionTLS12,
	}
	client = redis.NewClient(options)
}

func RedisGet(key string) (string, error) {
	ctx := context.TODO()
	value, err := client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", nil
		}
		return "", err
	}
	return value, nil
}

func RedisSet(key string, value string) error {
	ctx := context.TODO()
	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

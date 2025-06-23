package utils

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func InitRedis(addr string) {
	RedisClient = redis.NewClient(&redis.Options{Addr: addr})
}

// BlacklistToken stores a token ID until expiry
func BlacklistToken(jti string, ttl time.Duration) error {
	return RedisClient.Set(Ctx, "blacklist:"+jti, "true", ttl).Err()
}

func IsTokenBlacklisted(jti string) (bool, error) {
	res, err := RedisClient.Get(Ctx, "blacklist:"+jti).Result()
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return res == "true", nil
}

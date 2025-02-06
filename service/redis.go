package service

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RealRedisCaller struct {
	RedisClient redis.Client
}

func (r *RealRedisCaller) Get(ctx context.Context, key string) (int64, error) {
	return r.RedisClient.Get(ctx, key).Int64()
}

func (r *RealRedisCaller) Set(ctx context.Context, key string, value int64, expiration time.Duration) error {
	return r.RedisClient.Set(ctx, key, value, expiration).Err()
}

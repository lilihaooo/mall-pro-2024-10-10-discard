package redis

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

type ResString struct {
	Client *redis.Client
}

// 获取 Redis key 的值
func (r *ResString) GetValue(ctx context.Context, key string) (string, error) {
	val, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", nil
		}
		return "", err
	}
	return val, nil
}

// 设置 Redis key 的值
func (r *ResString) SetValue(ctx context.Context, key string, value string, expiration time.Duration) error {
	err := r.Client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

// 删除 Redis key
func (r *ResString) DeleteKey(ctx context.Context, key string) error {
	return r.Client.Del(ctx, key).Err()
}

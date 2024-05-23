/**
 * @Author: Nan
 * @Date: 2024/5/21 下午2:26
 */

package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"strconv"
	"sweet-cms/global"
	"sweet-cms/inter"
	"time"
)

type RedisUtil struct {
	client *redis.Client
}

func NewRedisUtil() *RedisUtil {
	return &RedisUtil{
		client: global.RedisClient,
	}
}

func withTimeout(duration time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), duration)
}

func (r *RedisUtil) Set(key string, value interface{}, expiration time.Duration) error {
	ctx, cancel := withTimeout(2 * time.Second)
	defer cancel()
	err := r.client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return inter.ErrCacheMiss
		}
		zap.S().Errorf("failed to set key %s: %v", key, err)
		return err
	}
	return nil
}

func (r *RedisUtil) Get(key string, value interface{}) error {
	ctx, cancel := withTimeout(2 * time.Second)
	defer cancel()
	val, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		zap.S().Errorf("failed to get key %s: %v", key, err)
		return err
	}
	switch v := value.(type) {
	case *string:
		*v = string(val)
	case *int:
		iv, err := strconv.Atoi(string(val))
		if err != nil {
			zap.S().Errorf("failed to convert string to int for key %s: %v", key, err)
			return err
		}
		*v = iv
	case *float64:
		fv, err := strconv.ParseFloat(string(val), 64)
		if err != nil {
			zap.S().Errorf("failed to convert string to float64 for key %s: %v", key, err)
			return err
		}
		*v = fv
	default:
		err := json.Unmarshal(val, value)
		if err != nil {
			zap.S().Errorf("failed to unmarshal value for key %s: %v", key, err)
			return err
		}
	}
	return nil
}

func (r *RedisUtil) Del(key string) error {
	ctx, cancel := withTimeout(2 * time.Second)
	defer cancel()
	err := r.client.Del(ctx, key).Err()
	if err != nil {
		zap.S().Errorf("failed to delete key %s: %v", key, err)
		return err
	}
	return nil
}

func (r *RedisUtil) Exists(keys ...string) (int64, error) {
	ctx, cancel := withTimeout(2 * time.Second)
	defer cancel()
	val, err := r.client.Exists(ctx, keys...).Result()
	if err != nil {
		return 0, fmt.Errorf("failed to check if keys exist: %v", err)
	}
	return val, nil
}

func (r *RedisUtil) Expire(key string, expiration time.Duration) (bool, error) {
	ctx, cancel := withTimeout(2 * time.Second)
	defer cancel()
	val, err := r.client.Expire(ctx, key, expiration).Result()
	if err != nil {
		return false, fmt.Errorf("failed to set expiration for key %s: %v", key, err)
	}
	return val, nil
}

func (r *RedisUtil) HSet(key, field string, value interface{}) error {
	ctx, cancel := withTimeout(2 * time.Second)
	defer cancel()
	err := r.client.HSet(ctx, key, field, value).Err()
	if err != nil {
		return fmt.Errorf("failed to hset key %s: %v", key, err)
	}
	return nil
}

func (r *RedisUtil) HGet(key, field string) (string, error) {
	ctx, cancel := withTimeout(2 * time.Second)
	defer cancel()
	val, err := r.client.HGet(ctx, key, field).Result()
	if err != nil {
		return "", fmt.Errorf("failed to hget key %s: %v", key, err)
	}
	return val, nil
}

func (r *RedisUtil) HDel(key string, fields ...string) (int64, error) {
	ctx, cancel := withTimeout(2 * time.Second)
	defer cancel()
	val, err := r.client.HDel(ctx, key, fields...).Result()
	if err != nil {
		return 0, fmt.Errorf("failed to hdel key %s: %v", key, err)
	}
	return val, nil
}

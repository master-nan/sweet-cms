/**
 * @Author: Nan
 * @Date: 2024/5/21 下午2:27
 */

package initialize

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"sweet-cms/config"
	"sweet-cms/global"
	"time"
)

func RedisClient() {
	conf := global.ServerConf.Redis
	dsn := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	options := &redis.Options{
		Addr:         dsn,
		Password:     conf.Password,
		DB:           conf.DB,
		PoolSize:     10,
		MinIdleConns: 5,
		MaxConnAge:   time.Hour,
	}

	client := redis.NewClient(options)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := client.Ping(ctx).Err()
	if err != nil {
		zap.S().Errorf("Failed to connect to Redis: %v", err)
		//panic(err)
	}
	global.RedisClient = client
	fmt.Println("Successfully connected to Redis")
}

func InitRedis(serverConfig *config.Server) (*redis.Client, error) {
	cfg := serverConfig.Redis
	client := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
		MaxConnAge:   time.Duration(cfg.MaxConnAge) * time.Second,
	})

	ctx := context.Background()
	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, err
	}
	return client, nil
}

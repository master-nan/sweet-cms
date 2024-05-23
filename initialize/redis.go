/**
 * @Author: Nan
 * @Date: 2024/5/21 下午2:27
 */

package initialize

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"sweet-cms/global"
	"time"
)

func RedisClient() {
	conf := global.ServerConf.Redis
	dsn := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	options := &redis.Options{
		Addr:     dsn,
		Password: conf.Password,
		DB:       conf.DB,
	}

	client := redis.NewClient(options)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := client.Ping(ctx).Err()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
		panic(err)
	}
	client.Options().PoolSize = 10
	client.Options().MinIdleConns = 5
	client.Options().MaxConnAge = time.Hour
	global.RedisClient = client
	fmt.Println("Successfully connected to Redis")
}

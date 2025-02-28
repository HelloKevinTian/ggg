package delayqueue

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(url string) *redis.Client {
	// redis://:tyhall51@192.168.20.119:8003/1
	if url == "" {
		url = "redis://:tyhall51@192.168.20.119:8003/1"
	}
	opt, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}

	// 以下均为go-redis的默认配置，无需显式设置
	// opt.PoolSize = 10 * runtime.GOMAXPROCS(0)
	// opt.MinIdleConns = 0
	// opt.DialTimeout = 5 * time.Second
	// opt.ReadTimeout = 3 * time.Second
	// opt.WriteTimeout = opt.ReadTimeout

	return redis.NewClient(opt)
}

// RunLuaScript 运行lua脚本
func RunLuaScript(ctx context.Context, cli *redis.Client, s string, keys []string, args ...interface{}) (interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	r := redis.NewScript(s)
	return r.Run(ctx, cli, keys, args...).Result()
}

package dao

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

const (
	// EmptyString ...
	EmptyString = ""
)

// RedisCache ...
type RedisCache struct {
	redis *redis.Client
}

// NewRedisClient ...
func NewRedisClient(master, password string, host []string, poolSize int) *RedisCache {
	client := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    master,
		SentinelAddrs: host,
		PoolSize:      poolSize,
		Password:      password,
	})
	return &RedisCache{redis: client}
}

// HGet ...
func (cache *RedisCache) HGet(key, field string) string {
	value, err := cache.redis.HGet(key, field).Result()
	if err != nil {
		return EmptyString
	}
	return value
}

// HGetAll ...
func (cache *RedisCache) HGetAll(key string) map[string]string {
	value, err := cache.redis.HGetAll(key).Result()
	if err != nil {
		return nil
	}
	return value
}

// HMSet ...
func (cache *RedisCache) HMSet(key string, fields map[string]interface{}) bool {
	if _, err := cache.redis.HMSet(key, fields).Result(); err != nil {
		return false
	}
	return true
}

// HSet ...
func (cache *RedisCache) HSet(key, field, value string) bool {
	if _, err := cache.redis.HSet(key, field, value).Result(); err != nil {
		return false
	}
	return true
}

// SAdd ...
func (cache *RedisCache) SAdd(key string, values []string) bool {
	if _, err := cache.redis.SAdd(key, values).Result(); err != nil {
		return false
	}
	return true
}

// SMembers ...
func (cache *RedisCache) SMembers(key string) []string {
	if values, err := cache.redis.SMembers(key).Result(); err == nil {
		return values
	}
	return nil
}

// Del ...
func (cache *RedisCache) Del(key ...string) bool {
	if _, err := cache.redis.Del(key...).Result(); err != nil {
		return false
	}
	return true
}

// IncrVersion ...
func (cache *RedisCache) IncrVersion(key string) (version int64, err error) {
	if version, err = cache.redis.Incr(key).Result(); err != nil {
		return
	}
	return
}

// GetIncrVersion ...
func (cache *RedisCache) GetIncrVersion(key string) (version int64, err error) {
	versionStr := EmptyString
	if versionStr, err = cache.redis.Get(key).Result(); err != nil && !errors.Is(err, redis.Nil) {
		return
	}
	if versionStr == EmptyString {
		return 0, nil
	}
	version, err = strconv.ParseInt(versionStr, 10, 64)
	return
}

// GetSet ...
func (cache *RedisCache) GetSet(key string) (values []string) {
	var index uint64
	var tempValues []string
	var err error
	for {
		tempValues, index, err = cache.redis.SScan(key, index, EmptyString, -1).Result()
		if err != nil {
			break
		}
		values = append(values, tempValues...)
		if index == 0 {
			break
		}
	}
	return
}

// Subscribe ...
func (cache *RedisCache) Subscribe(key string, fun func(message string)) {
	subscribe := cache.redis.Subscribe(key)
	for {
		msg := <-subscribe.Channel()
		fmt.Println(fmt.Printf("Subscribe Get Channel:[%s],Message:[%s]", msg.Channel, msg.Payload))
		go fun(msg.Payload)
	}
}

// Publish ...
func (cache *RedisCache) Publish(key, msg string) bool {
	if _, err := cache.redis.Publish(key, msg).Result(); err != nil {
		return false
	}
	return true
}

// RPop ...
func (cache *RedisCache) RPop(key string) (value string, isSuccess bool) {
	var err error
	if value, err = cache.redis.RPop(key).Result(); err != nil {
		return value, false
	}
	return value, true
}

// LPush ...
func (cache *RedisCache) LPush(key string, values []interface{}) (count int64, isSuccess bool) {
	var err error
	if count, err = cache.redis.LPush(key, values...).Result(); err != nil {
		return count, false
	}
	return count, true
}

// SetNX ...
func (cache *RedisCache) SetNX(key string, value interface{}, expiration time.Duration) (isSuccess bool) {
	isSuccess, err := cache.redis.SetNX(key, value, expiration).Result()
	if err != nil {
		return
	}
	return
}

// Expire ...
func (cache *RedisCache) Expire(key string, expiration time.Duration) (isSuccess bool) {
	isSuccess, err := cache.redis.Expire(key, expiration).Result()
	if err != nil {
		return
	}
	return
}

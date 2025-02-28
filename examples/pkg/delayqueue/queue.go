package delayqueue

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func ProduceTaskWithLua(rdb *redis.Client, queueName, taskID, taskData string, delay time.Duration) error {
	// 计算延时时间点
	execTime := time.Now().Add(delay).Unix()

	// Lua 脚本内容
	luaScript := `
        local zsetKey = KEYS[1]
        local hashKey = KEYS[2]
        local taskID = ARGV[1]
        local taskData = ARGV[2]
        local execTime = ARGV[3]

        -- 插入任务到 HASH
        redis.call('HSET', hashKey, taskID, taskData)

        -- 插入任务到 ZSET，分数为执行时间点
        redis.call('ZADD', zsetKey, execTime, taskID)

        return "OK"
    `

	// 执行 Lua 脚本
	keys := []string{
		fmt.Sprintf("%s:zset", queueName),  // ZSET 的键
		fmt.Sprintf("%s:tasks", queueName), // HASH 的键
	}
	args := []interface{}{
		taskID,   // 任务 ID
		taskData, // 任务内容
		execTime, // 执行时间戳
	}
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()
	_, err := RunLuaScript(ctx, rdb, luaScript, keys, args...)
	return err
}

func ConsumeTask(rdb *redis.Client, queueName string) (string, string, error) {
	// 当前时间戳
	now := time.Now().Unix()

	// Lua 脚本：原子性查找并移除任务
	luaScript := `
        local zsetKey = KEYS[1]
        local hashKey = KEYS[2]
        local currentTime = ARGV[1]

        -- 找到符合条件的任务
        local tasks = redis.call('ZRANGEBYSCORE', zsetKey, '-inf', currentTime, 'LIMIT', 0, 1)
        if #tasks == 0 then
            return nil
        end

        local taskID = tasks[1]

        -- 移除任务
        redis.call('ZREM', zsetKey, taskID)
        local taskData = redis.call('HGET', hashKey, taskID)
        redis.call('HDEL', hashKey, taskID)

        return {taskID, taskData}
    `

	// 执行脚本
	keys := []string{fmt.Sprintf("%s:zset", queueName), fmt.Sprintf("%s:tasks", queueName)}
	args := []interface{}{now}
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()
	result, err := RunLuaScript(ctx, rdb, luaScript, keys, args...)
	if err != nil {
		return "", "", err
	}

	// 解析返回值
	if result == nil {
		return "", "", nil
	}
	resArr := result.([]interface{})
	taskID := resArr[0].(string)
	taskData := resArr[1].(string)
	return taskID, taskData, nil
}

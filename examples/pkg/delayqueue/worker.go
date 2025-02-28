package delayqueue

import (
	"fmt"
	"ggg/utils"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestRedisDelayQueue() {
	rdb := NewRedisClient("")

	// 队列名称
	queueName := "delay_queue"

	// 消费
	go loopConsume(rdb, queueName)

	// 生产任务
	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(10 * time.Millisecond)
			i := i
			taskId := fmt.Sprintf("task%d", i)
			delay := time.Duration(utils.RandIn(5, 15)) * time.Second
			taskData := fmt.Sprintf(`Task Data %d delay %s`, i, delay)
			err := ProduceTaskWithLua(rdb, queueName, taskId, taskData, delay)
			if err != nil {
				fmt.Println("Failed to produce task:", err)
				return
			}
			fmt.Println("Task produced successfully! ", taskId, taskData, delay)
		}
	}()

	time.Sleep(1000 * time.Second)
}

func loopConsume(rdb *redis.Client, queueName string) {
	intervalTime := 200 * time.Millisecond
	// intervalTime := 1 * time.Second
	idleTimer := time.NewTimer(intervalTime)
	var ready bool = true
	var success int
	for {
		select {
		// case <-readych:
		// 	ready = true
		case <-idleTimer.C:
			fmt.Println("redis timer IN success: ", success)
			if ready {
				taskID, taskData, err := ConsumeTask(rdb, queueName)
				if err != nil {
					fmt.Println("Failed to consume task:", err)
					resetTimer(idleTimer, intervalTime)
					continue
				}
				success++
				fmt.Printf("Consumed task: ID=%s, Data=%s\n", taskID, taskData)
				resetTimer(idleTimer, intervalTime)
				// ready = false
			}
		}

	}
}

func resetTimer(idleTimer *time.Timer, intervalTime time.Duration) {
	// fmt.Println("resetTimer IN")
	// if !idleTimer.Stop() {
	// 	<-idleTimer.C
	// }
	fmt.Println("redis timer Reset\n")
	idleTimer.Reset(intervalTime)
}

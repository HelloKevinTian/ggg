package some

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 测试悲观锁 乐观锁
func TestLock() {
	c := time.Tick(time.Second * 2)
	for range c {
		// fmt.Println("Tick every 1 seconds")
		tk()
	}
}

func tk() {
	value1 := 0
	var value2 int32
	atomic.StoreInt32(&value2, 0)
	value3 := 0
	ws := sync.WaitGroup{}
	ws.Add(100000)
	var lo sync.Mutex

	for i := 0; i < 100000; i++ {
		go func() {
			defer ws.Done()
			value1++
			atomic.AddInt32(&value2, 1) // 乐观锁
			func() {                    // 悲观锁
				lo.Lock()
				defer lo.Unlock()
				value3++
			}()
		}()
	}

	ws.Wait()
	fmt.Println(value1, value2, value3)
}

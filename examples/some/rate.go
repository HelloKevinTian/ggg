package some

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/ratelimit"
	"golang.org/x/time/rate"
)

func TestRate() {
	// testRateAllow()
	testRateWait()
	// testUberRate()
}

func testRateAllow() {
	limiter := rate.NewLimiter(rate.Every(500*time.Millisecond), 5) //频率：500毫秒取一次token，并发量为5个

	for i := 0; i < 20; i++ {
		if i == 5 || i == 10 || i == 15 {
			time.Sleep(2000 * time.Millisecond) //取完5个后休眠2秒又会产生4个可用的token（2000 / 500 = 4）
		}
		ok := limiter.Allow()
		if ok {
			fmt.Println("allow ", i)
		} else {
			fmt.Println("deny ", i)
		}
	}
}

func testRateWait() {
	limiter := rate.NewLimiter(rate.Every(1000*time.Millisecond), 1)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	prev := time.Now()
	for i := 0; i < 20; i++ {
		err := limiter.WaitN(ctx, 1) //5个取完后开始每等1秒取1个token执行
		now := time.Now()
		if err == nil {
			fmt.Println("allow ", i, now.Sub(prev))
		} else {
			fmt.Println("deny ", i, err)
		}
		prev = now
	}
	fmt.Println("cancel")
}

func testUberRate() {
	rl := ratelimit.New(10)
	prev := time.Now()
	for i := 0; i < 10; i++ {
		now := rl.Take()
		fmt.Println(i, now.Sub(prev))
		prev = now
	}
}

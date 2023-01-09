package some

import (
	"context"
	"fmt"
	"time"
)

func TestContext() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// go handle(ctx, 500*time.Millisecond)  //超时时间大于处理时间，所以有足够时间来处理请求
	go handle(ctx, 1500*time.Millisecond) //超时时间小于处理时间，所以处理请求会被中断
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

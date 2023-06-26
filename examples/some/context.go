package some

import (
	"context"
	"fmt"
	"time"
)

// 不要将 Context 塞到结构体里。直接将 Context 类型作为函数的第一参数，而且一般都命名为 ctx。
// 不要向函数传入一个 nil 的 context，如果你实在不知道传什么，标准库给你准备好了一个 context：todo。
// 不要把本应该作为函数参数的类型塞到 context 中，context 存储的应该是一些共同的数据。例如：登陆的 session、cookie 等。
// 同一个 context 可能会被传递到多个 goroutine，别担心，context 是并发安全的。
func TestContext() {
	//测试超时
	ct1()

	fmt.Printf("\n--------------\n")

	//测试Value
	ct2()

	fmt.Printf("\n--------------\n")

	//测试cancel
	ct3()
}

// 【1】ctx父子协程的标准用法
func ct1() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// go handle(ctx, 500*time.Millisecond)  //超时时间大于处理时间，所以有足够时间来处理请求
	go handle(ctx, 1500*time.Millisecond) //超时时间小于处理时间，所以处理请求会被中断

	<-ctx.Done()
	fmt.Println("main", ctx.Err())
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

type VvStr string

// 【2】测试ctx的value取值顺序
func ct2() {
	var k1 VvStr = "111"
	var k2 VvStr = "222"
	var k3 VvStr = "333"
	valCtx := context.WithValue(context.Background(), k1, "aaa111")
	valCtx1 := context.WithValue(valCtx, k2, "aaa222")
	valCtx2 := context.WithValue(valCtx1, k3, "aaa333")
	valCtx3 := context.WithValue(valCtx1, k1, "aaa444")

	// value遍历顺序：自已=>父ctx=>父父ctx
	fmt.Println(valCtx.Value(k1), valCtx1.Value(k1), valCtx2.Value(k1))                    // aaa111 aaa111 aaa111 优先返回自身存储的value，如没有就用父ctx的value
	fmt.Println(valCtx2.Value(k1), valCtx2.Value(k2), valCtx2.Value(k3))                   // aaa111 aaa222 aaa333 优先返回自身存储的value
	fmt.Println(valCtx3.Value(k1), valCtx3.Value(k2), valCtx3.Value(k3), valCtx.Value(k1)) // aaa444 aaa222 <nil> aaa111 并不会覆盖父级的value 兄弟节点的value相互隔离
}

// 【3】测试ctx的cancel顺序
func ct3() {
	ctx1, c1 := context.WithCancel(context.Background())

	ctx2, c2 := context.WithCancel(ctx1)

	ctx3, c3 := context.WithCancel(ctx2)

	go func() {
		time.Sleep(time.Second)
		_ = c1
		_ = c3
		// defer c1()
		defer c2()
		// defer c3()
	}()

	for {
		select {
		case <-ctx1.Done():
			fmt.Println("ctx1 canceled ", ctx1.Err(), ctx2.Err(), ctx3.Err())
			return
		case <-ctx2.Done():
			fmt.Println("ctx2 canceled ", ctx2.Err(), ctx2.Err(), ctx3.Err())
			return
		case <-ctx3.Done():
			fmt.Println("ctx3 canceled ", ctx3.Err(), ctx2.Err(), ctx3.Err())
			return
		}
	}

}

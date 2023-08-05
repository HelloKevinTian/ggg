package examples

import (
	"context"
	"fmt"
	"ggg/examples/dao"
	"ggg/examples/helper"
	"ggg/examples/some"
	"ggg/examples/some/container"
	"ggg/examples/some/somesort"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

// 定义函数映射表
var funcs = map[string]interface{}{
	"test":            test,
	"testMD5":         testMD5,
	"testHTTPClient":  testHTTPClient,
	"testReader":      testReader,
	"testWrite":       testWrite,
	"testMongo":       testMongo,
	"testSlice":       testSlice,
	"testLogger":      testLogger,
	"testLeak":        testLeak,
	"testContext":     testContext,
	"TestShortId":     TestShortId,
	"TestFindWorkDay": TestFindWorkDay,

	"TestQuickSort":  somesort.TestQuickSort,
	"TestHeapSort":   somesort.TestHeapSort,
	"TestMergeSort":  somesort.TestMergeSort,
	"TestBubbleSort": somesort.TestBubbleSort,
	"TestInsertSort": somesort.TestInsertSort,

	"TestPriorityQueue": container.TestPriorityQueue,
	"TestStack":         container.TestStack,
	"TestCircleQueue":   container.TestCircleQueue,
	"TestTreeDFS":       container.TestTreeDFS,

	"TestHTTPServer":   some.TestHTTPServer,
	"TestGinServer":    some.TestGinServer,
	"TestBTree":        some.TestBTree,
	"TestEmbed":        some.TestEmbed,
	"TestSession":      some.TestSession,
	"TestReflect":      some.TestReflect,
	"InvokeRouter":     some.InvokeRouter,
	"TestViper":        some.TestViper,
	"TestLock":         some.TestLock,
	"BatchN":           some.BatchN,
	"SendRambler":      some.SendRambler,
	"TestSort":         some.TestSort,
	"TestRate":         some.TestRate,
	"TestbiSearch":     some.TestbiSearch,
	"TestBreaker":      some.TestBreaker,
	"StartGachaServer": some.StartGachaServer,
	"TestBatch":        some.TestBatch,
	"TestErrGroup":     some.TestErrGroup,
	"TestContext":      some.TestContext,
	"TestCIDR":         some.TestCIDR,
	"TestAggregator":   some.TestAggregator,
}

func RunExamples() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic:", err)
		}
	}()
	if len(os.Args) < 3 {
		panic(`no func to exec, refer to:

	- go run main.go examples testMD5
	- go run main.go examples testSlice
	- go run main.go examples testContext
	- go run main.go examples TestShortId
	- go run main.go examples TestBatch
	- go run main.go examples TestReflect
	- go run main.go examples testReader
	- go run main.go examples testWrite
`)
	}

	curFunc := os.Args[2]
	fmt.Printf("============[%s] IN===============\n\n", curFunc)
	helper.Call(funcs, curFunc)
	fmt.Printf("\n============[%s] OUT===============", curFunc)
}

//============================Test IN===============================

const (
	// East ...
	East int = iota + 1
	// South ...
	South
	// West ...
	West
	// North ...
	North
)

// VOData ...
type VOData struct {
	Message string `json:"message"`
	Data    string `json:"data"`
	Time    int64  `json:"time"`
}

// VO ...
type VO struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Data   VOData `json:"data"`
}

func TestFindWorkDay() {
	// var s, _ = time.Parse(time.RFC3339, "2021-12-26T16:00:00.000Z")
	// var e, _ = time.Parse(time.RFC3339, "2021-12-31T16:00:00.000Z")
	t1, _ := time.Parse(time.RFC3339, "2022-01-03T16:00:00.000Z")
	t2, _ := time.Parse(time.RFC3339, "2022-01-05T15:59:59.000Z")
	t3, _ := time.Parse(time.RFC3339, "2021-12-26T16:00:00.000Z")
	t4, _ := time.Parse(time.RFC3339, "2021-12-31T15:59:59.000Z")
	num := FindWorkDay(t1, t2)  //4,5,6
	num1 := FindWorkDay(t3, t4) //27 28 29 30 31
	fmt.Println(num, num1)
}

func FindWorkDay(s, e time.Time) (workDays int) {
	ctz, _ := time.LoadLocation("Asia/Shanghai")
	s = s.In(ctz)
	e = e.In(ctz)
	sWeek := s.Weekday()
	diffDay := int(e.Sub(s).Hours() / 24)
	fmt.Println(s, e, sWeek, diffDay)
	for i := int(sWeek); i <= int(sWeek)+diffDay; i++ {
		w := time.Weekday(i % 7)
		if w != time.Sunday && w != time.Saturday {
			workDays++
		}
	}
	return
}

// 测试context.Context
type str string

var cKey str = "name"

func testContext() {
	ctx, cancel := context.WithCancel(context.Background())
	go watchContext(ctx, "【监控1】")
	go watchContext(ctx, "【监控2】")
	go watchContext(ctx, "【监控3】")

	valueCtx := context.WithValue(ctx, cKey, "【监控K-V】")
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(ctx.Value(cKey), "监控退出，停止了...")
				return
			default:
				fmt.Println(ctx.Value(cKey), "goroutine 监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}(valueCtx)

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	time.Sleep(5 * time.Second)
}
func watchContext(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出，停止了...")
			return
		default:
			fmt.Println(name, "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}

func testLogger() {
	// helper.QuickStart()
	helper.TestLogger()
}

func testSlice() {
	slice := []int{10, 20, 30, 40, 50}

	newSlice := slice[1:3] //长度为2，容量为4

	fmt.Println("切片增长前:", slice, newSlice, len(newSlice), cap(newSlice))

	newSlice = append(newSlice, 88)
	newSlice = append(newSlice, 99)
	newSlice = append(newSlice, 100) //因为容量为4，此处已超出容量，所以容量扩大至2倍（新容量为8）

	fmt.Println("切片增长后:", slice, newSlice, len(newSlice), cap(newSlice))
}

func testMongo() {
	dao.TestMongo()
}

// 测试延时器、定时器、G内存泄漏
func testLeak() {
	// for range time.Tick(time.Second * 2) {
	// 	fmt.Println("Tick")
	// 	testMemLeak()
	// }

	// c := time.NewTicker(time.Second)
	// for {
	// 	<-c.C
	// 	fmt.Println("NewTicker")
	// }

	// t := time.NewTimer(time.Second * 2)
	// for range t.C {
	// 	fmt.Println("NewTimer")
	// 	t.Reset(time.Second * 1)
	// }
	tt := time.AfterFunc(time.Second*3, testMemLeak)
	defer tt.Stop()
	time.Sleep(time.Second * 5)
}
func testMemLeak() {
	num := 6
	for index := 0; index < num; index++ {
		resp, _ := http.Get("https://www.baidu.com")
		// if err != nil {
		// 	fmt.Println("Get error")
		// 	return
		// }
		// defer resp.Body.Close()
		_, _ = ioutil.ReadAll(resp.Body)
	}
	fmt.Printf("此时goroutine个数= %d\n", runtime.NumGoroutine())
}

// testWrite
func testWrite() {
	w := helper.NewChanWriter()

	go func() {
		defer w.Close()
		w.Write([]byte("Stream"))
		w.Write([]byte("Me"))
		w.Write([]byte("Yoo\n"))
	}()

	// for i := 0; i < 100; i++ {
	// 	v, ok := <-w.Chan()
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println(v)
	// }

	var byteArr []byte
	for c := range w.Chan() {
		fmt.Println(c)
		byteArr = append(byteArr, c)
	}
	os.Stdout.Write(byteArr)
}

// testReader ...
func testReader() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := helper.NewRot13Reader(s)
	io.Copy(os.Stdout, r)
}

// test ...
func test() {
	//测试G执行
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("welcome GG")
	}()
	wg.Wait()
	// time.Sleep(1 * time.Second)

	//测试struct初始化和引用
	a := VOData{Message: "1", Data: "2", Time: 1}
	b := &VOData{Message: "1", Data: "2", Time: 1}
	a.Message = "3"
	b.Message = "4"
	fmt.Println(West, North, East, South, a, *b)

	//闭包测试
	i := incr()
	b1 := i()
	b2 := i()
	fmt.Println("测试闭包函数", b1, b2, incr()())
}

// incr 闭包函数
func incr() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

// testMD5 测试md5
func testMD5() {
	md5 := helper.MD5("123")
	println("func:testMD5", md5)
}

// testHTTPClient 测试get、post请求和结果填充
func testHTTPClient() {
	// vo1 := &VO{}
	// vo2 := &VO{}
	var vo1 VO
	var vo2 VO
	getChan := make(chan bool)
	postChan := make(chan bool)

	go func(getChan chan<- bool) {
		getChan <- helper.GetJSON("https://auth.ftsview.com/getTime", &vo1)
	}(getChan)

	go func(postChan chan<- bool) {
		postChan <- helper.PostJSON("https://auth.ftsview.com/getTime", `{"a":1}`, &vo2)
	}(postChan)

	if ok1 := <-getChan; ok1 {
		fmt.Printf("\nget vo: %+v", vo1)
	}
	if ok2 := <-postChan; ok2 {
		fmt.Printf("\npost vo: %+v", vo2)
	}
}

func TestShortId() {
	fmt.Println(GenShortCode(6))
	fmt.Println(GenShortCode(6))
	fmt.Println(GenShortCode(6))
}

func GenShortCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

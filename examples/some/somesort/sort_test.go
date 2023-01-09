package somesort

import (
	"math/rand"
	"testing"
	"time"
)

func randomArr(count int) []int {
	rand.Seed(time.Now().Unix())
	var ret = []int{}
	for i := 0; i < count; i++ {
		v := rand.Int() % count
		ret = append(ret, v)
	}
	return ret
}

// var arr = randomArr(10000)

// go test -benchmem -run=^$ -bench ^Benchmark* gcargo/some/somesort
// 测试结果：数据量较小时，各排序算法差距不大，数据量较大时，快排、堆排序较优
// BenchmarkBubble100-8               72862             16295 ns/op            2040 B/op          8 allocs/op
// BenchmarkInsert100-8               85812             13762 ns/op            2040 B/op          8 allocs/op
// BenchmarkQuick100-8                90346             12963 ns/op            2040 B/op          8 allocs/op
// BenchmarkHeap100-8                 88894             13283 ns/op            2040 B/op          8 allocs/op
// BenchmarkBubble10000-8                14          78892623 ns/op          386297 B/op         20 allocs/op
// BenchmarkInsert10000-8                80          14824122 ns/op          386296 B/op         20 allocs/op
// BenchmarkQuick10000-8               1297            898842 ns/op          386296 B/op         20 allocs/op
// BenchmarkHeap10000-8                1159           1002628 ns/op          386297 B/op         20 allocs/op

func BenchmarkBubble100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := randomArr(100)
		bubbleSort(arr)
	}
}

func BenchmarkInsert100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := randomArr(100)
		insertSort(arr)
	}
}

func BenchmarkQuick100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := randomArr(100)
		quickSort(arr)
	}
}

func BenchmarkHeap100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := randomArr(100)
		heapSort(arr)
	}
}

func BenchmarkBubble10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := randomArr(10000)
		bubbleSort(arr)
	}
}

func BenchmarkInsert10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := randomArr(10000)
		insertSort(arr)
	}
}

func BenchmarkQuick10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := randomArr(10000)
		quickSort(arr)
	}
}

func BenchmarkHeap10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := randomArr(10000)
		heapSort(arr)
	}
}

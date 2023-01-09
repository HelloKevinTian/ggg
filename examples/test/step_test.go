// 测试程序局部性原理
// go test -benchmem -run=^$ -bench ^BenchmarkLoop* gcargo/test
//
// BenchmarkLoopStep1-8            1000000000               0.9521 ns/op          0 B/op          0 allocs/op
// BenchmarkLoopStep16-8           100000000               10.29 ns/op            0 B/op          0 allocs/op
// PASS
// ok      gcargo/test     2.100s
package test

import "testing"

func loopStep(nums []int, step int) {
	l := len(nums)
	for i := 0; i < step; i++ {
		for j := i; j < l; j += step {
			nums[j] = 4
		}
	}
}

func BenchmarkLoopStep1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loopStep([]int{}, 1)
	}
}

func BenchmarkLoopStep16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loopStep([]int{}, 16)
	}
}

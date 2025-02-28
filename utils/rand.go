package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// RandCode 随机n位数字
func RandCode(n int) string {
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		code = fmt.Sprintf("%s%d", code, rand.Intn(10))
	}
	return code
}

// RandIn 范围随机数 [min, max)
func RandIn(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

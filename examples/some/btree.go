package some

import (
	"fmt"
	"reflect"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func walk(t *Tree, ch chan int) {
	sendValue(t, ch)
	close(ch) //此处必须close ch，否则会引发39行的 <- ch 处死锁
}

func sendValue(t *Tree, ch chan int) {
	if t != nil {
		sendValue(t.Left, ch)
		ch <- t.Value
		sendValue(t.Right, ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func same(t1, t2 *Tree) bool {
	s1 := make([]int, 0)
	s2 := make([]int, 0)
	c1 := make(chan int)
	c2 := make(chan int)
	s1 = getArr(t1, c1)
	s2 = getArr(t2, c2)
	return reflect.DeepEqual(s1, s2)
}

func getArr(t *Tree, ch chan int) []int {
	ss := make([]int, 0)
	go walk(t, ch)

	for v := range ch {
		ss = append(ss, v)
	}
	return ss
}

// TestBTree ...
func TestBTree() {
	fmt.Println("-----TestBTree-----")
	var ss []int
	ch := make(chan int)
	t := NewTree(1)
	ss = getArr(t, ch)
	fmt.Println(ss)

	fmt.Println(same(NewTree(1), NewTree(2)))
}

package some

import (
	"fmt"
	"reflect"
	"sort"
)

type Asset struct {
	ID      int
	Click   int
	Install int
}

type objSlice []Asset

func (p objSlice) Len() int {
	return len(p)
}

func (p objSlice) Less(i, j int) bool {
	return p[i].Click > p[j].Click
}

func (p objSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func TestSort() {
	// t1()
	// t2()

	t3("Install")
	t3("Click")

	// t4()
}

// 语法糖
func t1() {
	il := []int{89, 14, 5, 17, 95, 11}
	sl := []string{"89", "14", "b5", "8", "17", "c", "95", "a11"}
	sort.Ints(il)
	sort.Strings(sl)
	fmt.Println(il, sl)
}

// 实现sort.Interface
func t2() {
	pl := objSlice{Asset{1, 100, 10}, Asset{2, 90, 100}, Asset{3, 300, 1000}}
	sort.Sort(pl)
	// r := sort.Reverse(pl)
	fmt.Println(pl)
}

//自定义排序方法
func t3(s string) {
	pl := []Asset{{1, 100, 110}, {2, 90, 2100}, {3, 300, 1000}}

	SortSlice(pl, s, true)
	fmt.Println(pl)
}

// 倒序 内置类型 IntSlice StringSlice Float64Slice
func t4() {
	il := []int{89, 14, 5, 17, 95, 11}
	sort.Sort(sort.Reverse(sort.IntSlice(il)))
	fmt.Println(il)
}

// SortSlice 结构体切片排序 使用方法如下：
//  type Asset struct {
//    ID int
//	  Click   int
//	  Install int
//  }
//  el := []Asset{{1, 100, 110}, {2, 90, 2100}, {3, 300, 1000}}
//  SortSlice(el, "Install", true)
func SortSlice(el interface{}, s string, desc bool) {
	sort.Slice(el, func(i, j int) bool {
		elLeft := reflect.ValueOf(el)
		eLeft := elLeft.Index(i)
		left := eLeft.FieldByName(s)

		elRight := reflect.ValueOf(el)
		eRight := elRight.Index(j)
		right := eRight.FieldByName(s)
		if desc {
			return left.Int() > right.Int()
		} else {
			return left.Int() < right.Int()
		}
	})
}

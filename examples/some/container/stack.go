package container

import "fmt"

type Stack struct {
	Items []string
	Count int
	n     int
}

func (s *Stack) Push(item string) {
	s.Items = append(s.Items, item)
	s.Count++
}

func (s *Stack) Pop() string {
	len := len(s.Items)
	v := s.Items[len-1]
	s.Items = s.Items[:len-1]
	s.Count--
	return v
}

func TestStack() {
	s := &Stack{
		Items: []string{},
	}
	s.Push("aaa")
	s.Push("bbb")
	s.Push("ccc")
	s.Push("ddd")
	fmt.Println(s.Items)
	s.Pop()
	s.Pop()
	fmt.Println(s.Items)

	fmt.Println("模拟浏览器前进后退")
	x := &Stack{
		Items: []string{},
	}
	y := &Stack{
		Items: []string{},
	}
	fmt.Println("访问a")
	x.Push("a")
	fmt.Println(x.Items, y.Items)

	fmt.Println("访问b")
	x.Push("b")
	fmt.Println(x.Items, y.Items)

	fmt.Println("访问c")
	x.Push("c")
	fmt.Println(x.Items, y.Items)

	fmt.Println("后退")
	y.Push(x.Pop())
	fmt.Println(x.Items, y.Items)

	fmt.Println("前进")
	x.Push(y.Pop())
	fmt.Println(x.Items, y.Items)
}

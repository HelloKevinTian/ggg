package container

import "fmt"

// 牺牲一个存储空间，head前面不存东西，当tail在head前就满了
type MyCircularQueue struct {
	items []int
	head  int
	tail  int
}

func NewCircularQueue(k int) *MyCircularQueue {
	return &MyCircularQueue{
		items: make([]int, k+1, k+1), //预留一个存储空间，用于区分队满和队空
		head:  0,
		tail:  0,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.items[this.tail] = value
	this.tail = (this.tail + 1) % cap(this.items)
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head + 1) % cap(this.items)
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.items[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	if this.tail == 0 {
		return this.items[cap(this.items)-1]
	}
	return this.items[this.tail-1]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.head == this.tail
}

func (this *MyCircularQueue) IsFull() bool {
	return this.head == (this.tail+1)%cap(this.items)
}

func TestCircleQueue() {
	circularQueue := NewCircularQueue(3) // 设置长度为 3
	r1 := circularQueue.EnQueue(1)       // 返回 true
	r2 := circularQueue.EnQueue(2)       // 返回 true
	r3 := circularQueue.EnQueue(3)       // 返回 true
	r4 := circularQueue.EnQueue(4)       // 返回 false，队列已满
	r5 := circularQueue.Rear()           // 返回 3
	r6 := circularQueue.IsFull()         // 返回 true
	r7 := circularQueue.DeQueue()        // 返回 true
	r8 := circularQueue.EnQueue(4)       // 返回 true
	r9 := circularQueue.Rear()           // 返回 4
	fmt.Println(r1, r2, r3, r4, r5, r6, r7, r8, r9)
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k)
 * param_1 := obj.EnQueue(value)
 * param_2 := obj.DeQueue()
 * param_3 := obj.Front()
 * param_4 := obj.Rear()
 * param_5 := obj.IsEmpty()
 * param_6 := obj.IsFull()
 */

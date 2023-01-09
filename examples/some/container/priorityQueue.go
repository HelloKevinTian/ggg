package container

import (
	"container/heap"
	"fmt"
)

//以下实现优先级队列
// This example demonstrates a priority queue built using the heap interface.
// Item 是 priorityQueue 中的元素
type Item struct {
	Key      string
	Priority int
	// index 是 Item 在 heap 中的索引号
	// Item 加入 Priority Queue 后， Priority 会变化时，很有用
	// 如果 Item.Priority 一直不变的话，可以删除 index
	index int
}

// PQ implements heap.Interface and holds entries.
type PQ []*Item

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push 往 pq 中放 Item
func (pq *PQ) Push(x interface{}) {
	temp := x.(*Item)
	temp.index = len(*pq)
	*pq = append(*pq, temp)
}

// Pop 从 pq 中取出最优先的 Item
func (pq *PQ) Pop() interface{} {
	temp := (*pq)[len(*pq)-1]
	temp.index = -1 // for safety
	*pq = (*pq)[0 : len(*pq)-1]
	return temp
}

// update modifies the priority and value of an Item in the queue.
func (pq *PQ) update(item *Item, value string, priority int) {
	item.Key = value
	item.Priority = priority
	heap.Fix(pq, item.index)
}

func TestPriorityQueue() {
	h := &PQ{}
	heap.Init(h)
	heap.Push(h, &Item{
		Key:      "joe",
		Priority: 5,
	})
	heap.Push(h, &Item{
		Key:      "trace",
		Priority: 3,
	})
	heap.Push(h, &Item{
		Key:      "lucy",
		Priority: 8,
	})
	e := &Item{
		Key:      "lily",
		Priority: 2,
	}
	heap.Push(h, e)
	heap.Push(h, &Item{
		Key:      "kev",
		Priority: 6,
	})
	// heap.Remove(h, 3)
	fmt.Println("------ ", e)
	h.update(e, "gogo", 9)

	fmt.Println(heap.Pop(h))
	fmt.Println(heap.Pop(h))
	fmt.Println(heap.Pop(h))
	fmt.Println(heap.Pop(h))
	fmt.Println(heap.Pop(h))
}

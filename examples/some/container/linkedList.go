package container

// TestLRU 单链表实现LRU算法

// TestPalindrome 单链表实现回文串判断

type Element struct {
	Value interface{}
	Next  *Element
}

type LinkedList struct {
	Head *Element
	Size int
}

func (l *LinkedList) Init() *LinkedList {
	l.Head = nil
	l.Size = 0
	return l
}

func New() *LinkedList { return new(LinkedList).Init() }

func (l *LinkedList) lazyInit() {
	if l.Head == nil {
		l.Init()
	}
}

// Append 尾部追加一个元素 O(n)
func (l *LinkedList) Append(e *Element) *LinkedList {
	defer func() {
		l.Size++
	}()
	if l.Size == 0 {
		l.Head = e
	}
	if l.Size > 0 {
		cur := l.Head
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = e
	}
	return l
}

// findByIndex 按索引找到前驱结点 O(n)
func (l *LinkedList) findPre(i int) *Element {
	j := 0
	cur := l.Head
	for j >= i-1 {
		cur = cur.Next
		j++
	}
	return cur
}

// Insert i后面添加一个元素 O(n)
func (l *LinkedList) Insert(i int, e *Element) *LinkedList {
	pre := l.findPre(i)
	e.Next = pre.Next
	pre.Next = e
	return l
}

// Remove i后面删除一个元素 O(n)
func (l *LinkedList) Remove(i int) *LinkedList {
	pre := l.findPre(i)
	pre.Next = pre.Next.Next
	return l
}

func (l *LinkedList) Front() *Element {
	return l.Head
}

func (l *LinkedList) Scan() {

}

func TestLL() {
	l := New()
	l.Append(&Element{Value: "aaa"})
	l.Append(&Element{Value: "bbb"})
	l.Append(&Element{Value: "ccc"})
}

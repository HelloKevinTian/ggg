package container

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestTreeDFS() {
	//    3
	//    ^
	//  9   20
	//       ^
	//     15  7
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	t := buildTree(preorder, inorder)
	fmt.Println(t, PreOrder(t), InOrder(t), PostOrder(t), LayerOrder(t), levelOrder(t))
	//[3 9 20 15 7] [9 3 15 20 7] [9 15 7 20 3] [3 9 20 15 7] [[3] [9 20] [15 7]]
}

//已知前序和中序 序列，求二叉树结构
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	idx := 0
	for idx = 0; idx < len(inorder)-1; idx++ {
		if preorder[0] == inorder[idx] {
			break
		}
	}
	if idx > 0 {
		root.Left = buildTree(preorder[1:idx+1], inorder[0:idx])
	}
	if idx < len(inorder)-1 {
		root.Right = buildTree(preorder[idx+1:], inorder[idx+1:])
	}
	return root
}

//先序遍历
func PreOrder(t *TreeNode) []int {
	preorder := []int{}
	preorder = append(preorder, t.Val)
	if t.Left != nil {
		preorder = append(preorder, PreOrder(t.Left)...)
	}
	if t.Right != nil {
		preorder = append(preorder, PreOrder(t.Right)...)
	}
	return preorder
}

//中序遍历
func InOrder(t *TreeNode) []int {
	inorder := []int{}
	if t.Left != nil {
		inorder = append(inorder, InOrder(t.Left)...)
	}
	inorder = append(inorder, t.Val)
	if t.Right != nil {
		inorder = append(inorder, InOrder(t.Right)...)
	}
	return inorder
}

//后续遍历
func PostOrder(t *TreeNode) []int {
	postorder := []int{}
	if t.Left != nil {
		postorder = append(postorder, PostOrder(t.Left)...)
	}
	if t.Right != nil {
		postorder = append(postorder, PostOrder(t.Right)...)
	}
	postorder = append(postorder, t.Val)
	return postorder
}

//层序遍历 输出一维数组
func LayerOrder(t *TreeNode) []int {
	layerorder := []int{}
	if t == nil {
		return layerorder
	}
	q := new(Queue)
	q.Push(t)

	for q.Len() != 0 {
		root := q.Pop().(*TreeNode)
		layerorder = append(layerorder, root.Val)
		if root.Left != nil {
			q.Push(root.Left)
		}
		if root.Right != nil {
			q.Push(root.Right)
		}
	}
	return layerorder
}

//层序遍历 输出二维数组
func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	q := new(Queue)
	q.Push(root)

	for q.Len() != 0 {
		list := []int{}
		count := q.Len()
		for count > 0 {
			node := q.Pop().(*TreeNode)
			list = append(list, node.Val)

			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
			count--
		}
		result = append(result, list)
	}
	return result
}

type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Pop() interface{} {
	if len(*q) == 0 {
		return nil
	}
	first := (*q)[0]
	*q = (*q)[1:]
	return first
}

func (q *Queue) Len() int {
	return len(*q)
}

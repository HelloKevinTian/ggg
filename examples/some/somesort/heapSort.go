package somesort

import "fmt"

func TestHeapSort() {
	arr := []int{5, 10, 1, 9, 2, 11, 5, 8, 7, 3}
	fmt.Println(heapSort(arr))
}

//堆排序的基本思想是：将待排序序列构造成一个大顶堆，此时，整个序列的最大值就是堆顶的根节点。将其与末尾元素进行交换，此时末尾就为最大值。然后将剩余n-1个元素重新构造成一个堆，这样会得到n个元素的次小值。如此反复执行，便能得到一个有序序列了

// 每个节点值都 >= 其左右孩子的值称为大顶堆，每个节点值都 <= 其左右孩子的值称为小顶堆
// 堆映射为数组后 左右孩子结点为 2i+1 和 2i+2
// 大顶堆：arr[i] >= arr[2i+1] && arr[i] >= arr[2i+2]
// 小顶堆：arr[i] <= arr[2i+1] && arr[i] <= arr[2i+2]
// 第一个非叶子节点一定是：len(arr)/2 - 1
func heapSort(arr []int) []int {
	//构建大顶堆
	for i := len(arr)/2 - 1; i >= 0; i-- {
		adjustHeap(arr, i, len(arr))
	}

	//重新调整结构+交换堆顶和末尾元素
	for j := len(arr) - 1; j >= 0; j-- {
		swap1(arr, j, 0)
		adjustHeap(arr, 0, j)
	}

	return arr
}

func adjustHeap(arr []int, i, length int) {
	temp := arr[i]                              //先取出当前元素i
	for k := i*2 + 1; k < length; k = k*2 + 1 { //从i结点的左子结点开始，也就是2i+1处开始
		if k+1 < length && arr[k] < arr[k+1] { //如果左子结点小于右子结点，k指向右子结点
			k++
		}
		if arr[k] > temp { //如果子节点大于父节点，将子节点值赋给父节点（不用进行交换）
			arr[i] = arr[k]
			i = k
		} else {
			break
		}
	}

	arr[i] = temp //将temp值放到最终的位置
}

func swap1(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

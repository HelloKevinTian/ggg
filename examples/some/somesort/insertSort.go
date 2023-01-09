package somesort

import (
	"fmt"
)

func TestInsertSort() {
	arr := []int{5, 10, 1, 9, 2, 11, 5, 8, 7, 3}
	bubbleSort(arr)
	fmt.Println(arr)
}

// 顺序插入排序，改进的可以进行二分插入
// 插入排序的原理是，从第二个数开始向右侧遍历，每次均把该位置的元素移动至左侧，放在一个正确的位置（比左侧大，比右侧小）
func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] { //发生逆序，往前插入
			temp := arr[i]
			j := i - 1
			for j >= 0 && arr[j] > temp {
				arr[j+1] = arr[j]
				j--
			}
			arr[j+1] = temp
		}
	}
}

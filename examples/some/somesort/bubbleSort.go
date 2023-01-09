package somesort

import "fmt"

var count1 int
var count2 int

func TestBubbleSort() {
	arr1 := []int{5, 10, 1, 9, 2, 11, 5, 8, 7, 3}
	bubbleSort(arr1)
	fmt.Println(arr1, len(arr1), count1)

	arr2 := []int{5, 10, 1, 9, 2, 11, 5, 8, 7, 3}
	bubbleSortBetter(arr2)
	fmt.Println(arr2, len(arr2), count2)
}

// 普通版本
func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[i] > arr[j] {
				count1++
				swap(arr, i, j)
			}
		}
	}
}

// 优化版本 若在一趟中没有发生逆序，则该序列已有序
func bubbleSortBetter(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		exchange := true
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				count2++
				swap(arr, j, j+1)
				exchange = false
			}
		}
		if exchange {
			break
		}
	}
}

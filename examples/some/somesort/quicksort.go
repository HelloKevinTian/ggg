package somesort

import "fmt"

func TestQuickSort() {
	arr := []int{5, 10, 1, 9, 2, 11, 5, 8, 7, 3}
	quickSort(arr)
	fmt.Println(arr)
}

// 快速排序的原理是，首先找到一个数pivot把数组‘平均’分成两组，使其中一组的所有数字均大于另一组中的数字，此时pivot在数组中的位置就是它正确的位置。然后，对这两组数组再次进行这种操作
func quickSort(arr []int) {
	_quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr []int, left, right int) {
	if left < right {
		pivot := partition(arr, left, right)
		_quickSort(arr, left, pivot-1)
		_quickSort(arr, pivot+1, right)
	}
}

// func partition(arr []int, left, right int) int {
// 	pivot := left
// 	index := pivot + 1

// 	for i := index; i <= right; i++ {
// 		if arr[i] < arr[pivot] {
// 			swap(arr, i, index)
// 			index++
// 		}
// 	}
// 	swap(arr, pivot, index-1)
// 	return index - 1
// }
func partition(arr []int, left int, right int) int {
	for left < right {
		for left < right && arr[left] <= arr[right] {
			right--
		}
		if left < right {
			swap(arr, left, right)
			left++
		}

		for left < right && arr[left] <= arr[right] {
			left++
		}
		if left < right {
			swap(arr, left, right)
			right--
		}
	}
	return left
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

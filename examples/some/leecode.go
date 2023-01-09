package some

import "fmt"

//二分法查找
func biSearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		middle := (left + right) / 2
		if target < arr[middle] {
			right = middle - 1
		} else if target > arr[middle] {
			left = middle + 1
		} else {
			return middle
		}
	}
	return -1
}

func TestbiSearch() {
	var arr = []int{1, 2, 3, 5, 7, 9, 10, 15, 18}
	fmt.Println("二分查找结果为: ", biSearch(arr, 9), biSearch(arr, 1), biSearch(arr, 18))
}

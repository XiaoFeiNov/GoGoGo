package algorithm

import (
	"sort"
)

// 递归实现二分查找
func BinarySearchByRecursion(arr []int, value int) int {
	if len(arr) == 0 {
		return -1
	}

	sort.Slice(arr, func(i, j int) bool { // 二分查找法需要在有序的条件下进行
		return i < j
	})
	return searchValue(arr, value, 0, len(arr) - 1)
}

func searchValue(arr []int, value, low, high int) int {
	if low > high { // low一定小于high
		return -1
	}

	mid :=  low + (high - low) / 2 // 此处注意需要加上low，low不一定是0所以需要加上已排除的数量
	if value < arr[mid] {
		return searchValue(arr, value, low, mid - 1)
	} else if value > arr[mid] {
		return searchValue(arr, value, mid + 1, high)
	} else {
		 return mid
	}
	return -1
}

// 循环条件实现二分查找
func BinarySearch(arr []int, value int) int {
	if len(arr) == 0 {
		return -1
	}

	sort.Slice(arr, func(i, j int) bool { // 二分查找法需要在有序的条件下进行
		return i < j
	})
	low, high := 0, len(arr) - 1
	for low <= high {
		mid := low + (high - low) / 2
		if value < arr[mid] {
			high = mid - 1
		} else if value > arr[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
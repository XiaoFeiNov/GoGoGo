package mySort

import "fmt"

func BubbleSort(arr []int)  {
	length := len(arr)
	for i := 0; i < length - 1; i++ {
		for j := 0; j < length - 1 - i; j++ { // 从零开始，所以和j + 1比，所以条件是< length - 1 - i
			if arr[j] > arr[j + 1] {
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
			}
		}
	}
}

func InsertionSort(arr []int)  {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i - 1] {
			tmp := arr[i]
			j := i
			for true {
				if j <= 0 || tmp > arr[j - 1] {
					break
				}
				arr[j] = arr[j - 1]
				j--
			}
			arr[j] = tmp
		}
	}
	arr = append(arr, 666)
	// 切片是一个包含数组指针、容量int、长度int的结构体，当修改数组内容时会改变原有切片，但超过容量时，不会改变原切片
	fmt.Println("--slice_append--", arr)
}

func BinarySearch(arr []int, x int) int {
	low, high := 0, len(arr) - 1
	for low <= high {
		mid := low + (high - low) / 2
		if arr[mid] == x {
			return mid
		} else if arr[mid] > x {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func ShellSort(arr []int)  {
	length := len(arr)
	for gap := length / 2; gap > 0; gap = gap / 2  {
		for i := gap; i < length ; i++ {
			j := i
			for j >= gap && arr[j] < arr[j - gap] {
				arr[j], arr[j - gap] = arr[j - gap], arr[j]
				j -= gap
			}
		}
	}
}

func QuickSort(arr []int, left, right int) {
	if left < right {
		pivot := partitioning(arr, left, right)
		QuickSort(arr, left, pivot - 1)
		QuickSort(arr, pivot + 1, right)
	}
}

func partitioning(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1
	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}
	arr[pivot], arr[index - 1] = arr[index - 1], arr[pivot]
	return index - 1
}

func MergeSort(a []int) []int {
	n := len(a)
	if n < 2 {
		return a
	}
	mid := n / 2
	left := a[0:mid]
	right := a[mid:]
	return merge(MergeSort(left), MergeSort(right))
}

func merge(left, right []int) (result []int) {
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	for len(left) > 0 {
		result = append(result, left[0])
		left = left[1:]
	}
	for len(right) > 0 {
		result = append(result, right[0])
		right = right[1:]
	}
	return
}
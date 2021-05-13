package mySort

// 冒泡、插入、选择排序平均时间复杂度都是n^2
// 冒泡排序，稳定的排序算法，适合n较小的情况
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n - 1; i++ {
		for j := 0; j < n - 1 - i; j++ {
			if arr[j] > arr[j + 1] {
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
			}
		}
	}
}

// 插入排序，稳定的排序算法，适合n较小的情况
func insertSort(arr []int) {
	n := len(arr)
	for i := 0; i < n - 1; i++ {
		for j := i + 1; j > 0; j-- {
			if arr[j] < arr[j - 1] {
				arr[j], arr[j - 1] = arr[j - 1], arr[j]
			}
		}
	}
}

// 选择排序，不稳定的排序算法（如：A 50, B 50, C 100，排序后为CBA），适合n较小的情况
func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n - 1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			arr[min], arr[i] = arr[i], arr[min]
		}
	}
}

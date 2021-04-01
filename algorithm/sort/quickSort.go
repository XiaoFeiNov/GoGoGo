package sort

func QuickSort(arr []int)  {
	if len(arr) < 2 {
		return
	}
	sort(arr, 0, len(arr) - 1)
}

func sort(arr []int, left, right int)  {
	if left < right {
		pivot := partitioning(arr, left, right) // 找出中心点，左边的都比中心点小，右边的都比中心点大
		sort(arr, left, pivot - 1) // 找出[0, pivot)之间的中心点
		sort(arr, pivot + 1, right) // 找出(pivot, right]之间的中心点
	}
}

func partitioning(arr []int, left, right int)(int) {
	pivot := left
	index := pivot + 1
	for i := index; i <= right; i++  {
		if arr[i] < arr[pivot] {
			arr[index], arr[i] = arr[i], arr[index]
			index++
		}
	}
	arr[pivot], arr[index - 1] = arr[index - 1], arr[pivot]
	return index - 1
}
package algorithm

// reversalArray，将数组元素顺序颠倒
func reversalArray(arr []int){
	n := len(arr)
	for i := 0; i < n / 2; i++ {
		arr[i], arr[n - 1 - i] = arr[n - 1 - i], arr[i]
	}
}

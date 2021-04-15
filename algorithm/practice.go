package algorithm

// reversalArray，将数组元素顺序颠倒
func reversalArray(arr []int){
	n := len(arr)
	for i := 0; i < n / 2; i++ {
		arr[i], arr[n - 1 - i] = arr[n - 1 - i], arr[i]
	}
}

// 斐波那契数列，0 1 1 2 3 ...
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x + y
	}
	return x
}

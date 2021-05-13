package algorithm

import (
	"fmt"
	"time"
)

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

func twoSum(a []int) (count int) {
	n := len(a)
	start := time.Now().UnixNano()
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if a[i] + a[j] == 0 {
				count++
			}
		}
	}
	time.Sleep(1 * time.Second)
	fmt.Printf("program spent: %d ns\n", time.Now().UnixNano() - start - 1000000000)
	return
}

func twoSumFast(a []int) (count int) {
	n := len(a)
	start := time.Now().UnixNano()
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if a[i] + a[j] == 0 {
				count++
			}
		}
	}
	fmt.Printf("program spent: %d ns\n", time.Now().UnixNano() - start - 1000000000)
	return
}
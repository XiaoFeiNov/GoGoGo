package algorithm

import (
	"fmt"
	"testing"
)

func TestReversalArray(t *testing.T)  {
	arr := []int{0, 1, 2, 3, 4, 5}
	fmt.Println("before reversal array = ", arr)
	reversalArray(arr)
	fmt.Println("after reversal array = ", arr)
}

func TestFib(t *testing.T)  {
	i := 2
	result := fib(i)
	fmt.Printf("fib(%d) = %d\n", i, result)
}

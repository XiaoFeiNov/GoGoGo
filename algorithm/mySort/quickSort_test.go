package mySort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T)  {
	arr := []int{15, 12, 5, 10, 8, 2}
	QuickSort(arr)
	fmt.Println("arr：", arr)
}
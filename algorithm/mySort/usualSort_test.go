package mySort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	a := []int{ 10, 2, 6, 5, 3}
	bubbleSort(a)
	fmt.Println("a = ", a)
	b := []int{ 10, 2, 6, 5, 3}
	insertSort(b)
	fmt.Println("b = ", b)
	c := []int{ 10, 2, 6, 5, 3}
	selectionSort(c)
	fmt.Println("c = ", c)
}

package algorithm

import (
	"fmt"
	"testing"
)

func TestBinarySearchByRecursion(t *testing.T)  {
	index := BinarySearchByRecursion([]int{0, 1, 2, 3, 4, 6}, 5)
	fmt.Println("BinarySearchByRecursion find result = ", index)
}

func TestBinarySearch(t *testing.T)  {
	index := BinarySearch([]int{0, 1, 2, 3, 4, 6}, 5)
	fmt.Println("BinarySearch find result = ", index)
}

package main

import "fmt"

func main()  {
	a := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println("a: ", a)
	reverseSlice(a)
	fmt.Println("reverseSlice: ", a)
	reverseSlice2(a)
	fmt.Println("reverseSlice2: ", a)
}

func reverseSlice(a []int) {
	for i := 0; i < len(a) / 2; i++ {
		a[i], a[len(a) - 1 - i] = a[len(a) - 1 - i], a[i]
	}
}

func reverseSlice2(a []int) {
	for i, j := 0, len(a) - 1; i < j; i, j = i + 1, j - 1 {
		a[i], a[j] = a[j], a[i]
	}
}
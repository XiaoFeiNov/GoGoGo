package main

import "fmt"

const (
	CONST_A = iota
	CONST_B
	CONST_C = 666
	CONST_D			// 与上面一致，666
	CONST_E = iota	// 计数器前面记了4位，所以这里是4
	CONST_F
)

const (
	CONST_BIT_A = 1 << iota
	CONST_BIT_B
	CONST_BIT_C = 666
	CONST_BIT_D				 	// 与上面一致，666
	CONST_BIT_E = 1 << iota		// 计数器前面记了4位，所以这里是1 * (2 ^ 4)
	CONST_BIT_F
)

func main()  {
	fmt.Printf("CONST_A = %d\n", CONST_A)
	fmt.Printf("CONST_B = %d\n", CONST_B)
	fmt.Printf("CONST_C = %d\n", CONST_C)
	fmt.Printf("CONST_D = %d\n", CONST_D)
	fmt.Printf("CONST_E = %d\n", CONST_E)
	fmt.Printf("CONST_F = %d\n", CONST_F)

	fmt.Printf("CONST_BIT_A = %d\n", CONST_BIT_A)
	fmt.Printf("CONST_BIT_B = %d\n", CONST_BIT_B)
	fmt.Printf("CONST_BIT_C = %d\n", CONST_BIT_C)
	fmt.Printf("CONST_BIT_D = %d\n", CONST_BIT_D)
	fmt.Printf("CONST_BIT_E = %d\n", CONST_BIT_E)
	fmt.Printf("CONST_BIT_F = %d\n", CONST_BIT_F)
}

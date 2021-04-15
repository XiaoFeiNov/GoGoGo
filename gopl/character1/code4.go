package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type selfInt int

func main()  {
	var i selfInt = 1
	j := 2
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(int(i)))
	fmt.Println(int(i) == j)
	unitTransform() // go run code4.go 5
}

func unitTransform()  {
	args := os.Args[1 : ]
	for _, arg := range args {
		num, _ := strconv.Atoi(arg)
		fmt.Printf("%d 斤 = %d g", num, num * 500)
	}
}
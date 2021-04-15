package main

import (
	"flag"
	"fmt"
)

var n = flag.Int("n", 0, "number")
var b = flag.Bool("b", false, "bool")
var s string

func main()  { // 执行命令行"go run code3.go -n=123 -b=true -s=abc"，用=进行赋值
	flag.StringVar(&s, "s", "aaa", "string")
	flag.Parse()
	fmt.Println("n = ", *n)
	fmt.Println("b = ", *b)
	fmt.Println("s = ", s)
}

package main

import (
	"fmt"
	"gogogo/gopl/myTest"
)

func main()  {
	test := new(myTest.TestStr)
	test.Name = "test"
	test.Num = 10
	test.SetSize(10) // TestStr的size是小写，包外无法获取相当于私有变量，只能通过方法对该成员进行操作
	fmt.Println("test: ", test)
	fmt.Println("test.GetSize: ", test.GetSize())
}

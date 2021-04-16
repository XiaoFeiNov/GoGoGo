package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Info struct {
	Name	string
	Age		int		`json:"age"` // 通过Tag对字段名字进行设置
	Sex		string	`json:"sex"`
	Class 	int
}

func main()  {
	studentA := Info{
		Name: 	"tom",
		Age: 	20,
		Sex:  	"man",
		Class:	6,
	}
	jsonInfo, err := json.Marshal(studentA)
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}
	fmt.Printf("jsonInfo: %s\n", jsonInfo)

	aInfo := new(Info)
	if err := json.Unmarshal(jsonInfo, aInfo); err != nil { // json.Unmarshal第二个变量需要填指针变量
		fmt.Println("err: ", err)
		os.Exit(2)
	}
	fmt.Printf("aInfo: %v\n", aInfo) // %#v，打印结构体，包括字段名称

	var name struct{Name string} // 通过设置，反序列化部分字段
	if err := json.Unmarshal(jsonInfo, &name); err != nil { // json.Unmarshal第二个变量需要填指针变量
		fmt.Println("err: ", err)
		os.Exit(2)
	}
	fmt.Printf("name: %v\n", name) // %#v，打印结构体，包括字段名称
}

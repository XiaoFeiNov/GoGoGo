package codeTest

import (
	"fmt"
	"io"
	"os"
)

// io.ReadFull：从Reader中读取len(buf)长度字节，与io.Read不同之处在于ReadFull必须准确读取len(buf)数量字节，多了少了都会返回错误
// 读完了n个字节后，Reader里面就少了n个字节（更正：不是少了字节，而是每次读取指针都会移动）
func readFull() {
	buff := make([]byte, 40, 40) // 一个汉字为三个字节，一个中文标点符号两个字节
	file, _ := os.Open("./data.txt")
	if n1, err1 := io.ReadFull(file, buff); err1 != nil {
		fmt.Println("err1 = ", err1, n1)
	}
	fmt.Println("buff = ", string(buff))
	fmt.Println(file.Seek(0, io.SeekCurrent))
	buff2 := make([]byte, 40, 40)
	if _, err2 := io.ReadFull(file, buff2); err2 != nil {
		fmt.Println("err2 = ", err2)
	}
	fmt.Println("buff2 = ", string(buff2))
	fmt.Println(file.Seek(0, io.SeekCurrent))
}

// https://studygolang.com/topics/4332?fr=sidebar
func read() {
	bytes := make([]byte, 10)
	file, _ := os.Open("./data1.txt")
	fmt.Println("seek__0")
	fmt.Println(file.Seek(0, io.SeekCurrent))
	file.Read(bytes)
	fmt.Println("seek__1")
	fmt.Println(file.Seek(0, io.SeekCurrent))

	fmt.Println(1, bytes)
	fmt.Println(2, string(bytes))

	file.Write([]byte("hello"))
	fmt.Println("seek__2")
	fmt.Println(file.Seek(0, io.SeekCurrent))
	file.Sync()
	fmt.Println("seek__3")
	fmt.Println(file.Seek(0, io.SeekCurrent))
	bytes2 := make([]byte, 10)
	_, err := file.Read(bytes2)
	fmt.Println("seek__4")
	fmt.Println(file.Seek(0, io.SeekCurrent))
	fmt.Println("err: ", err)
	fmt.Println(3, bytes2)
	fmt.Println(4, string(bytes2))
}

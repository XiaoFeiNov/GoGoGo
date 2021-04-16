package main

import (
	"bytes"
	"fmt"
)

func main()  {
	value := []int{0, 1, 2, 3}
	s1 := intToStringByByte(value)
	fmt.Println("s1 = ", s1)
	s2 := intToStringByBuf(value)
	fmt.Println("s2 = ", s2)
}

func intToStringByByte(value []int) string {
	var result []byte
	result = append(result, '[')
	for i, v := range value {
		if i > 0 {
			result = append(result, ',')
		}
		result = append(result, byte(v))
	}
	result = append(result, ']')
	return string(result)
}


func intToStringByBuf(value []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range value {
		if i > 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	//buf.WriteRune(']')
	buf.WriteByte(']')
	// 当向bytes.Buffer添加任意字符的UTF8编码时，最好使用bytes.Buffer的WriteRune方法
	// 但是，WriteByte方法对于写入类似'['和']'等ASCII字符则会更加有效
	return buf.String()
}

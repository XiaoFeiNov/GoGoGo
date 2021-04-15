package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("Args length：", len(os.Args))
	fmt.Println("Args：", os.Args) // os.Args[0]是命令本身的名字，参数是从下标1开始即os.Args[1:]

	// args填url:  (https协议不能漏掉)
	url := os.Args[1]
	fmt.Println("url = ", url)
	ch := make(chan string)

	for i := 0; i < 2; i++ {
		go func(i int) {
			start := time.Now()
			fmt.Printf("%d----------------------------%d", i, i)
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("http.Get err = %v", err)
			}
			// 读取body
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("ioutil.ReadAll err = %v", err)
			}
			fmt.Printf("body = %s", body)

			// 拷贝body，输出拷贝的字节长度
			// 如果上面先读取了body，这里进行拷贝就会数据长度为零，原因猜测是ReadAll导致指针位移到已读的位置，所以拷贝的时候拷贝的是空的
			//bodyLength, err := io.Copy(ioutil.Discard, resp.Body)
			//if err != nil {
			//	fmt.Printf("io.Copy err = %v", err)
			//}
			//fmt.Printf("bodyLength = %d", bodyLength)

			// 关闭resp的body流，防止leak resource
			resp.Body.Close()
			ch <- string(body)
			fmt.Println("time.since : ", time.Since(start).Seconds())
		}(i)
	}

	for i := 0; i < 2; i++ {
		<- ch
	}
}

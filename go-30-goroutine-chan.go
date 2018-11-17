/**
Channel是连接并行协程(goroutine)的通道。
你可以向一个通道写入数据然后从另外一个通道读取数据。

 */
package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	go func() {
		messages <- "ping"
	}()
	go func() {
		messages <- "pong"
	}()

	msg1 := <- messages
	fmt.Println(msg1)
	msg2 := <- messages
	fmt.Println(msg2)
}
/**

超时对那些连接外部资源的程序来说是很重要的，否则就需要限定执行时间。
在Go里面实现超时很简单。
我们可以使用channel和select很容易地做到。

 */
package main

import (
	"time"
	"fmt"
)

func main() {

	c1 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	// 这里使用select来实现超时，`res := <-c1`等待通道结果，
	// `<- Time.After`则在等待1秒后返回一个值，因为select首先
	// 执行那些不再阻塞的case，所以这里会执行超时程序，如果
	// `res := <-c1`超过1秒没有执行的话

	select {
	case res := <- c1:
		fmt.Println(res)
	case <- time.After(time.Second * 1):
		fmt.Println("timeout 1 sec")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	select {
	case res := <- c2:
		fmt.Println(res)
	case <- time.After(time.Second * 3):
		fmt.Println("timeout 2 sec")
	}

}


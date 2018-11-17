/**
关闭通道的意思是该通道将不再允许写入数据。
这个方法可以让通道数据的接受端知道数据已经全部发送完成了。
 */

package main

import "fmt"

func main() {
	//
	jobs := make(chan int, 5)
	done := make(chan bool)

	// 这里是数据接收端协程，它重复使用`j, more := <-jobs`来从通道
	// jobs获取数据，这里的more在通道关闭且通道中不再有数据可以接收的
	// 时候为false，我们通过判断more来决定所有的数据是否已经接收完成。
	// 如果所有数据接收完成，那么向done通道写入true
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}

		}
	}()

	for j := 1; j<= 3;j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	close(jobs)

	fmt.Println("sent all jobs")

	fmt.Println(<- done)
}

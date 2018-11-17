/**

频率控制是控制资源利用和保证服务高质量的重要机制。
Go可以使用goroutine，channel和ticker来以优雅的方式支持频率控制。
 */

package main

import (
	"time"
	"fmt"
)

func main() {
	requests := make(chan int, 5)
	for i:=1;i<=5;i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 200)

	// 通过阻塞从limiter通道接受数据，我们将请求处理控制在每隔200毫秒
	// 处理一个请求，注意`<-limiter`的阻塞作用。
	for req := range requests {
		<- limiter
		fmt.Println("request:", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	// 填充burstyLimiter，先发送3个数据
	for i:= 0;i < 3;i ++ {
		burstyLimiter <- time.Now()
	}

	// 然后每隔200毫秒再向burstyLimiter发送一个数据，这里是不断地
	// 每隔200毫秒向burstyLimiter发送数据
	go func() {
		for t:= range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	// 这里模拟5个请求，burstyRequests的前面3个请求会连续被处理，
	// 因为burstyLimiter被先连续发送3个数据的的缘故，而后面两个
	// 则每隔200毫秒处理一次
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

}

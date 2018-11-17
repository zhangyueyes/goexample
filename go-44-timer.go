package main

import (
	"time"
	"fmt"
)

func main() {
	timer1 := time.NewTimer(time.Second*2)

	// 这里`<-timer1.C`在timer的通道`C`上面阻塞等待，直到有个值发送给该
	// 通道，通知通道计时器已经等待完成。
	// timer.NewTimer方法获取的timer1的结构体定义为
	// type Ticket struct{
	//  C <-chan Time
	//}
	<- timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<- timer2.C
		fmt.Println("Timer 2 expired")
	}()

	//time.Sleep(time.Second * 2)
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

}


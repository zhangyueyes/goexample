/**
Timer是让你等待一段时间然后去做一件事情，这件事情只会做一次。

而Ticker是让你按照一定的时间间隔循环往复地做一件事情，除非你手动停止它。

 */

package main

import (
	"time"
	"fmt"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)

	go func() {
		for t:= range ticker.C {
			fmt.Println("Tick at:", t)
		}
	}()

	time.Sleep(time.Millisecond * 2001)

	ticker.Stop()

	fmt.Println("Ticker stopped")
}
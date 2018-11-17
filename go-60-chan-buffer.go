/**
默认情况下，通道是不带缓冲区的。
发送端发送数据，同时必须又接收端相应的接收数据。
而带缓冲区的通道则允许发送端的数据发送和接收端的数据获取处于异步状态，就是说发送端发送的数据可以放在缓冲区里面，可以等待接收端去获取数据，而不是立刻需要接收端去获取数据。
不过由于缓冲区的大小是有限的，所以还是必须有接收端来接收数据的，否则缓冲区一满，数据发送端就无法再发送数据了。
 */

package main

import "fmt"

func main() {
	message := make(chan string, 2)

	message <- "buffered"
	message <- "channel"

	fmt.Println(<- message)
	fmt.Println(<- message)

	message <- "channel"
	message <- "channel"
	fmt.Println(<- message)
	fmt.Println(<- message)
}

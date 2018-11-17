/**
for循环是Go语言唯一的循环结构。这里有三个基本的for循环类型。
 */

package main

import "fmt"

func main() {
	i := 1
	for i<=4 {
		fmt.Println(i)
		i++
	}

	for j:=7; j<=9;j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("Loop")
		break
	}
}


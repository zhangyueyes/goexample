package main

import "fmt"

func main() {
	//
	add10 := closure(10)
	fmt.Println(add10(5))
	fmt.Println(add10(6))

	add20 := closure(20)
	fmt.Println(add20(5))
}


func closure(x int) func(y int) int  {
	return func(y int) int {
		return x + y
	}
}

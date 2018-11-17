package main

import "fmt"

func main() {
	//
	result := adder()

	for i:=0; i< 10; i++ {
		fmt.Println(i, result(i))
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

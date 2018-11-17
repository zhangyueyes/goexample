package main

import "fmt"

type Callback func(x, y int) int

func test1(x, y int, callback Callback) int {
	return callback(x,y)
}

func add(x, y int) int {
	return x + y
}

func main()  {
	x, y := 1,4
	fmt.Println(test1(x,y,add))
}

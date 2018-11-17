package main

import "fmt"

func test2(s []int) []int {
	s = append(s, 3)
	return s
}

func main() {
	s1 := make([]int, 0)
	s1 = test2(s1)

	fmt.Println(s1)
}


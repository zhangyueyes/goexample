package main

import (
	"fmt"
	"math"
)

const s = "my name is zhangyue"

func main() {
	//
	fmt.Println(s)

	const n = 20000.1

	const m = 40000 / n

	fmt.Println(float32(m))

	fmt.Println(float64(m))

	fmt.Println(math.Sin(n))

}

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))

	fmt.Println(rand.Float64())

	// 这个方法可以用来生成其他数值范围内的随机数，
	// 例如`5.0 <= f < 10.0`
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// 为了使随机数生成器具有确定性，可以给它一个seed
	s1 := rand.NewSource(42)
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// 如果源使用一个和上面相同的seed，将生成一样的随机数
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
}

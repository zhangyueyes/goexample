package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// &操作符用来取得i变量的地址
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// 指针类型也可以输出
	fmt.Println("pointer:", &i)
}

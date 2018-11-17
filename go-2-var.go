package main

import "fmt"

/*
	var  vs   :=
	bool
	string
	int int8 int16 int32 int64
	uint uint8 uint16 uint32 uint64 uintptr
	byte // uint8 的别名
	rune // int32 的别名 代表一个Unicode码
	float32 float64
	complex64 complex128
*/

func main() {
	//
	var a = "initial"

	fmt.Println(a)

	var b, c = 2, 3
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "other"
	fmt.Println(f)

}

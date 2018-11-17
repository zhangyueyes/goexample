package main

import "fmt"

/*
	数组是一个具有相同数据类型的元素组成的固定长度的有序集合。
	在Go语言中，数组是值类型，长度是类型的组成部分，也就是说"[10]int"和“[20]int”是完全不同的两种数组类型。
	同类型的两个数组支持"=="和"!="比较，但是不能比较大小。
	数组作为参数时，函数内部不改变数组内部的值，除非是传入数组的指针。
	数组的指针：*[3]int
	指针数组：[2]*int

*/
func main() {
	//
	var a [5]int
	fmt.Println("emp", a)

	var b [5]string
	fmt.Println("emp", b)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get", a[4])

	fmt.Println("len:", len(a))

	c := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", c)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

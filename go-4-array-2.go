package main

import "fmt"

/**
拥有固定长度是数组的一个特点，但是这个特点有时候会带来很多不便，尤其在一个集合元素个数不固定的情况下。
这个时候我们更多地使用切片

可以用new创建数组，并返回数组的指针
 */

func main()  {
	//
	var a= new([5]int)

	test(a)

	fmt.Println(a, len(a))

}

func test(a *[5]int){
	a[1] = 5
}

/**

Go支持匿名函数，匿名函数可以形成闭包。
闭包函数可以访问定义闭包的函数定义的内部变量

 */


package main

import "fmt"

func intSeq() func() int  {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main()  {
	//
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
	fmt.Println(nextInt())
}

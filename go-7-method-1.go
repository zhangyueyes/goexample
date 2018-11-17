package main

import "fmt"

/**
一般的函数定义叫做函数，定义在结构体上面的函数叫做该结构体的方法

方法的定义限定类型可以为结构体类型
也可以是结构体指针类型
区别在于如果限定类型是结构体指针类型
那么在该方法内部可以修改结构体成员信息

 */

type rect struct {
	width, height int
}

func (r *rect) area() int  {
	return r.width * r.height
}

func (r rect) perim() int {
	r.width = 30
	return 2*r.width + 2*r.height
}

func main()  {
	//
	r := rect{width:10, height:5}

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	rr := r

	fmt.Println("area-rr:", rr.area())
	fmt.Println("perim-rr:", rr.perim())

	fmt.Println("area-rp:", rp.area())
	fmt.Println("perim-rp:", rp.perim())
}
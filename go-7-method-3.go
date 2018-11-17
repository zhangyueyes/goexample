
/**
使用匿名字段，实现模拟继承。
即可直接访问匿名字段（匿名类型或匿名指针类型）的方法这种行为类似“继承”。
访问匿名字段方法时，有隐藏规则，这样我们可以实现override效果。

 */

package main

import "fmt"

type Person1 struct {
	Id int
	Name string
}

type Student struct {
	Person1
	Score int
}

func main()  {
	//
	p := Student{Person1{2, "zhangyue"}, 25}
	p.test()
}

func (p Person1) test()  {
	fmt.Println("person test")
}

func (p Student) test()  {
	fmt.Println("student test")
}


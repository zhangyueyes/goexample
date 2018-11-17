/**

instance.method(args)->(type).func(instance,args)

为了区别这两种方式，官方文档中将
左边的称为Method Value，
右边则是Method Expression。
Method Value是包装后的状态对象，
总是与特定的对象实例关联在一起（类似闭包，拐带私奔），
而Method Expression函数将Receiver作为第一个显式参数，
调用时需额外传递。

注意：对于Method Expression，T仅拥有T Receiver方法，*T拥有（T+*T）所有方法。

 */

package main

import "fmt"

type Person struct {
	Id int
	Name string
}

func (P Person) test(x int) {
	fmt.Println("Id:", P.Id, "Name:", P.Name)
	fmt.Println("x:", x)
}

func main()  {
	p := Person{1, "张跃"}

	p.test(2)

	var f1 func(int) = p.test
	f1(3)

	Person.test(p, 4)

	var f2 func(Person, int) = Person.test
	f2(p, 5)
}


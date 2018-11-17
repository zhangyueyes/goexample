
/**
如果一个函数的参数是接口类型，那么我们可以使用命名接口来调用这个函数
比如这里的正方形square和圆形circle都实现了接口geometry，
那么它们都可以作为这个参数为geometry类型的函数的参数。
在measure函数内部，Go知道调用哪个结构体实现的接口方法。

也就是说如果结构体A实现了接口B定义的所有方法，那么A也是B类型。
 */

package main

import (
	"math"
	"fmt"
)

type geometry interface {
	area() float64
	perim() float64
}

type square struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (s square) area() float64  {
	return s.width * s.height
}

func (s square) perim() float64  {
	return 2 * s.width + 2*s.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main()  {
	s := square{width:3, height:4}
	c := circle{radius:5}
	// 这里circle和square都实现了geometry接口，所以
	// circle类型变量和square类型变量都可以作为measure
	// 函数的参数
	measure(s)
	measure(c)
}




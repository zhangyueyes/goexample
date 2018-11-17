package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}


func main() {
	p:= point{1,2}

	// 下面的%v打印了一个point结构体的对象的值
	fmt.Printf("%v\n", p)

	// 如果所格式化的值是一个结构体对象，那么`%+v`的格式化输出
	// 将包括结构体的成员名称和值
	fmt.Printf("%+v\n", p)

	// `%#v`格式化输出将输出一个值的Go语法表示方式。
	fmt.Printf("%#v\n", p)

	fmt.Printf("%T\n", p)
	fmt.Printf("%t\n", p)

	fmt.Printf("%d\n", 33)
	fmt.Printf("%b\n", 33)
	fmt.Printf("%x\n", 33)

	fmt.Printf("%c\n", 108)

	fmt.Printf("%f\n", 123.4)
	fmt.Printf("%s\n", "\"string\"")

	// `%x`以16进制输出字符串，每个字符串的字节用两个字符输出
	fmt.Printf("%x\n", "hex this")

	// 使用`%p`输出一个指针的值
	fmt.Printf("%p\n", &p)

	// 当输出数字的时候，经常需要去控制输出的宽度和精度。
	// 可以使用一个位于%后面的数字来控制输出的宽度，默认
	// 情况下输出是右对齐的，左边加上空格
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	// 你也可以指定浮点数的输出宽度，同时你还可以指定浮点数
	// 的输出精度
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// To left-justify, use the `-` flag.
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// 你也可以指定输出字符串的宽度来保证它们输出对齐。默认
	// 情况下，输出是右对齐的
	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	// 为了使用左对齐你可以在宽度之前加上`-`号
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// `Printf`函数的输出是输出到命令行`os.Stdout`的，你
	// 可以用`Sprintf`来将格式化后的字符串赋值给一个变量
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// 你也可以使用`Fprintf`来将格式化后的值输出到`io.Writers`
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}



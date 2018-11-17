/**
strings 标准库提供了很多字符串操作相关的函数

 */

package main

import "fmt"
import ss "strings"

// 这里给fmt.Println起个别名，因为下面我们会多处使用。
var p = fmt.Println

func main()  {

	p("Contains: ", ss.Contains("test", "es"))
	p("Count: ", ss.Count("test", "t"))
	p("HasPrefix: ", ss.HasPrefix("test", "te"))
	p("HasSuffix: ", ss.HasSuffix("test", "st"))
	p("Index: ", ss.Index("test", "e"))
	p("Join: ",ss.Join([]string{"a", "b"}, "-"))
	p("Repeat: ", ss.Repeat("a", 10))
	p("Replace: ", ss.Replace("fooooo", "o", "A", 3))
	p("Replace: ", ss.Replace("fooooo", "o", "A", 2))
	p("Replace: ", ss.Replace("fooooo", "o", "A", 1))
	p("Split: ", ss.Split("a-b-c-d", "-"))
	p("ToLower: ", ss.ToLower("TEST"))
	p("ToUpper: ", ss.ToUpper("test"))

	p()

	p("Len: ", len("zhangyue"))
	p("Char: ", string("hello"[2]))
}



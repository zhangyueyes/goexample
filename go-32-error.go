/**

在Go里面通常采用显式返回错误代码的方式来进行错误处理。
这个和Java或者Ruby里面使用异常或者是C里面运行正常返回结果，发生错误返回错误代码的方式不同。

Go的这种错误处理的方式使得我们能够很容易看出哪些函数可能返回错误，
并且能够像调用那些没有错误返回的函数一样调用。

 */

package main

import (
	"github.com/kataras/iris/core/errors"
	"fmt"
)

type argError struct {
	arg int
	prob string
}

// 你可以通过实现error接口的方法Error()来自定义错误
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New使用错误信息作为参数，构建一个基本的错误
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

func f2(arg int) (int, error)  {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}

	return arg + 3, nil
}


func main() {
	for _, i := range []int{7, 42} {
		if r,e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int {7, 42} {
		if r,e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// 如果你需要使用自定义错误类型返回的错误数据，你需要使用类型断言
	// 来获得一个自定义错误类型的实例才行。

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
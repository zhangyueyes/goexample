/**
Line Filters翻译一下大概是行数据过滤器。
简单一点就是一个程序从标准输入stdin读取数据，然后处理一下，将处理的结果输出到标准输出stdout。
grep和sed就是常见的行数据过滤器。
 */

package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

func main() {

	// 使用缓冲scanner来包裹无缓冲的`os.Stdin`可以让我们
	// 方便地使用`Scan`方法，这个方法会将scanner定位到下
	// 一行的位置
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())

		fmt.Println(ucl)
	}

	if err:= scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
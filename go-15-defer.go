/**
* Defer 用来保证一个函数调用会在程序执行的最后被调用。
*
* 通常用于资源清理工作。
**/

package main

import (
	"os"
	"fmt"
)

func main() {
	f := createFile("/tmp/defer.txt")

	defer closeFile(f)

	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}

	return f
}

func writeFile(f *os.File)  {
	fmt.Println("writing")
	fmt.Fprintln(f, "data-string")
}

func closeFile(f *os.File)  {
	fmt.Println("closing")
	f.Close()
}



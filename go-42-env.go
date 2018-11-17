package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	os.Setenv("NAME", "zhangyue")
	fmt.Println(os.Getenv("NAME"))
	fmt.Println(os.Getenv("name"))
	fmt.Println(os.Getenv("N"))

	fmt.Println()
	for _,e := range  os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0], "=", pair[1])
	}
}


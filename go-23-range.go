/**
range函数是个神奇而有趣的内置函数，你可以使用它来遍历数组，切片和字典。
当用于遍历数组和切片的时候，range函数返回索引和元素；
当用于遍历字典的时候，range函数返回字典的键和值。
 */

package main

import "fmt"

func main() {
	nums := []int{2,3,4}

	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum: ", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a":"apple", "b":"banana"}
	for k,v := range  kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range函数用来遍历字符串时，返回Unicode代码点。
	// 第一个返回值是每个字符的起始字节的索引，第二个是字符代码点，
	// 因为Go的字符串是由字节组成的，多个字节组成一个rune类型字符。
	for i, c := range "gopher" {
		fmt.Println(i,string(c))
	}
}





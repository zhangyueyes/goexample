package main

import (
	"fmt"
	"encoding/json"
)

/*
字典是Go语言内置的关联数据类型。
因为数组是索引对应数组元素，而字典是键对应值

make(map[键类型]值类型)

 */
func main()  {
	//
	m := make(map[string]int)

	m["k1"] = 9
	m["k2"] = 10

	fmt.Println("map:",m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)
	fmt.Println("len:", len(m))

	_, ok := m["k2"]
	fmt.Println("ok:", ok)

	l := map[string]string{"middle":"beautiful"}

	n := map[string]interface{}{"age":15,"name":map[string]string{}}
	n["name"] = l
	fmt.Println("map:", n)

	fmt.Println("age:", n["age"])

	data, err := json.Marshal(n)
	fmt.Println("error:",err)
	fmt.Println("data", data)

	p := map[string]interface{}{}
	//var p map[string]interface{}
	//p = make(map[string]interface{})
	p["name"] = "zhangyue"
	p["age"] = 28
	fmt.Println("map:",p)
}

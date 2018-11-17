package main

import "fmt"

type Person2 struct {
	name string
	age int
}

func main()  {
	fmt.Println(Person2{"Bob", 20})

	fmt.Println(Person2{name:"Bob", age:20})

	fmt.Println(Person2{name:"Bob"})

	fmt.Println(&Person2{name:"alic", age:30})

	s := Person2{name:"Lisa", age:50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.name)

	sp.age = 51
	fmt.Println(sp.age)

	fmt.Println(s.age)
}



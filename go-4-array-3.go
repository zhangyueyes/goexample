package main

import "fmt"

type User struct {
	Id int
	Name string
}

func main()  {
	//

	var darr []string

	darr = append(darr, "one")
	darr = append(darr, "two")
	darr = append(darr, "three")

	fmt.Println(darr)
	fmt.Println(len(darr))

	a := [...]User{
		{0, "User0"},
		{8, "User8"},
	}

	b := [...]*User{
		{0, "User0"},
		{8, "User8"},
	}

	fmt.Println(a, len(a))
	fmt.Println(b, len(b))

	var c []User

	c = append(c, User{9, "User9"})
	fmt.Println(c, len(c))

	c = append(c, User{10, "User10"})
	fmt.Println(c, len(c))
}

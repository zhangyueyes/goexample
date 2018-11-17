package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp: ", s)

	s[0] = "str_a"
	s[1] = "str_b"
	s[2] = "str_c"

	fmt.Println("set: ", s)
	fmt.Println("get: ", s[1])
	fmt.Println("len: ", len(s))

	s = append(s, "str_d")
	s = append(s, "str_e", "str_f")

	fmt.Println("apd: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)

	l := s[2:4]
	fmt.Println("sl1: ", l)

	l = s[:5]
	fmt.Println("sl2: ", l)

	l = s[2:]
	fmt.Println("sl3: ", l)

	t := []string{"a", "b", "c"}
	fmt.Println("dcl: ", t)

	twoD := make([][]int, 3)
	for i := 0; i< 3; i++ {
		innerLen := i +1
		twoD[i] = make([]int, innerLen)
		for j:=0;j<innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD: ", twoD)
}


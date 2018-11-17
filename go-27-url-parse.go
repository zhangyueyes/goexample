package main

import (
	"net/url"
	"fmt"
)

func main() {
	s := "postgres://user:pass@host.com:8888/path/to/api?q=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)
	fmt.Println(u.Host)
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	fmt.Println(u.User.Password())
	fmt.Println(u.Path)
	fmt.Println(u.RawQuery)
	fmt.Println(u.RequestURI())
	fmt.Println(u.Fragment)
	fmt.Println(u.Query())
}

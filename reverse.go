package main

import "fmt"

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}

func StrRev(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

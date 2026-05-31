package main

import "fmt"

func main() {
	n1 := 5
	result1 := InvertSign(&n1)
	fmt.Println(*result1)

	n2 := "hello"
	ResetString(&n2)
}

func InvertSign(x *int) *int {
	n1 := *x * 2
	n2 := *x - n1
	return &n2
}

func ResetString(s *string) {
	*s = ""
	fmt.Println(*s)
}



package main

import (
	"fmt"
)

func main() {
	var a int = 13
	var b int = 7
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a, b)
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a, b)

	a = a - b
	b = b + a
	a = b - a
	fmt.Println(a, b)
	a = a - b
	b = b + a
	a = b - a
	fmt.Println(a, b)
}

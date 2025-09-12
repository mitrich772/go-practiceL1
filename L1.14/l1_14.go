package main

import (
	"fmt"
	"reflect"
)

func identify(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int: ", v)
	case string:
		fmt.Println("string: ", v)
	case bool:
		fmt.Println("bool: ", v)
	default:
		if reflect.TypeOf(v).Kind() == reflect.Chan {
			fmt.Println("chan")
		} else {
			fmt.Println("Мутный тип")
		}
	}
}
func main() {
	a := 42
	b := "hello"
	c := true

	chInt := make(chan int)
	chString := make(chan string)
	chBool := make(chan bool)

	identify(a)
	identify(b)
	identify(c)
	identify(chInt)
	identify(chString)
	identify(chBool)

	var slice = []int{1, 2, 3}
	identify(slice)
}

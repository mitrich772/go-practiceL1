package main

import (
	"fmt"
	"strings"
)

var justString string

func createHugeString(n int) string {
	return strings.Repeat("Ы", n)
}

func someFunc() {
	v := createHugeString(1 << 10)       //2^10
	justString = string([]rune(v)[:100]) //Копируем в новый кусок памяти под символы
}

func main() {
	someFunc()
	fmt.Print(justString)
}

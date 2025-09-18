package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)

	fmt.Print("Введите a: ")
	fmt.Scan(a) // считываем a

	fmt.Print("Введите b: ")
	fmt.Scan(b) // считываем b

	// Сложение
	sum := new(big.Int).Add(a, b)
	fmt.Println("Сумма:", sum)

	// Вычитание
	diff := new(big.Int).Sub(a, b)
	fmt.Println("Разность:", diff)

	// Умножение
	prod := new(big.Int).Mul(a, b)
	fmt.Println("Произведение:", prod)

	// Деление целочисленное
	quot := new(big.Int).Div(a, b)
	fmt.Println("Частное:", quot)
}

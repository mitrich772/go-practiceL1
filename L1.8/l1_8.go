package main

import (
	"fmt"
)

func main() {
	var num int64
	fmt.Print("Напишите число: ")
	fmt.Scanln(&num)
	var i uint
	var bit int
	for {
		fmt.Printf("Исходное число: %d | %b\n", num, num)
		fmt.Print("Введите номер бита и значение (0 или 1): ")
		fmt.Scanln(&i, &bit)

		if bit != 0 {
			num |= 1<<i + 1 // i+1 из запримера Пример: для числа 5 (0101₂) установка 1-го бита в 0 даст 4 (0100₂).
		} else {
			num &^= 1<<i + 1
		}

		fmt.Printf("Итоговое число: %d | %b\n\n", num, num)
	}
}

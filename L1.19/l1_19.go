package main

import "fmt"

func swap(mas []rune, firstPos int, secondPos int) {
	t := mas[firstPos]
	mas[firstPos] = mas[secondPos]
	mas[secondPos] = t
}
func reverse(str string) string {
	runes := []rune(str)
	for i := 0; i < (len(runes)-1)-i; i++ {
		swap(runes, i, (len(runes)-1)-i)
	}
	return string(runes)
}
func main() {
	fmt.Println(reverse("главрыба"))
}
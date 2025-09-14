package main

import (
	"fmt"
)

func swap(mas []rune, firstPos int, secondPos int) {
	t := mas[firstPos]
	mas[firstPos] = mas[secondPos]
	mas[secondPos] = t
}
func reverse(runes []rune, from int, to int) []rune {
	for left, right := from, to; left < right; left, right = left+1, right-1 {
		swap(runes, left, right)
	}
	return runes
}
func reverseWords(toReverse string) string{
	sentense := []rune(toReverse)
	reverse(sentense, 0, len(sentense) - 1)
	lastSpace := -1
	for curRunePos, r := range sentense{
		if(r == ' '){
			reverse(sentense, lastSpace + 1, curRunePos - 1)
			lastSpace = curRunePos
		}
	}
	reverse(sentense, lastSpace+1, len(sentense)-1)
	return string(sentense)

}
func main() {
	fmt.Println(reverseWords("snow dog sun pupok arbyz"))
}
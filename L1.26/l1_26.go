package main

import (
	"fmt"
	"strings"
)

func checkUniq(s string) bool {
	runes := []rune(strings.ToLower(s))
	repMap := make(map[rune]bool) //Все кидаем в map смотрим потом в нем же повтор
	for _, r := range runes {
		if repMap[r] {
			return false
		}
		repMap[r] = true
	}
	return true
}

func main() {
	fmt.Println(checkUniq("Прривет"))
}

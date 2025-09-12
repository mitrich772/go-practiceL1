package main

import (
	"fmt"
)

type void struct{}

// Можно использовать то что уже написаноо в l1_11:)
func main() {
	animals := []string{"cat", "cat", "dog", "cat", "tree"}
	uniqueAnimals := make(map[string]void) //Ключи у map уникланы
	for _, v := range animals {
		uniqueAnimals[v] = void{}
	}
	for v := range uniqueAnimals {
		fmt.Println(v)
	}
}

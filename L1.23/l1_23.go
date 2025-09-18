package main

import "fmt"

func remove[T any](sl []T, i int) []T {
	copy(sl[i:], sl[i+1:])
	sl[len(sl)-1] = *new(T)
	return sl[:len(sl)-1]
}

func main() {
	sl := []int{1, 2, 3, 4, 5}
	fmt.Println(remove(sl,1))
}

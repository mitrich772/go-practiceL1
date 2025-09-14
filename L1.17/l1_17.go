package main

import "fmt"

func binarySearch(sl []int, f int) int {
	ptr := sl
	offset := 0

	for len(ptr) > 0 {
		currentPos := len(ptr) / 2

		if ptr[currentPos] == f {
			return offset + currentPos //угадали

		} else if ptr[currentPos] < f { // перелет
			offset += currentPos + 1
			ptr = ptr[currentPos+1:]

		} else { //недолет
			ptr = ptr[:currentPos]
		}
	}

	return -1
}

func main() {
	sl := []int{1, 2, 3, 4, 5, 7, 8, 9}
	fmt.Println(binarySearch(sl, 3))
}

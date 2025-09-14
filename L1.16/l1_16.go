package main

import "fmt"

func swap(mas []int, firstPos int, secondPos int) {
	t := mas[firstPos]
	mas[firstPos] = mas[secondPos]
	mas[secondPos] = t
}
func raskidAroundPivot(mas []int, left int, right int) int {
	wall := left - 1
	pivot := mas[right] //опорный элемент
	for current := left; current < right; current++ {
		if mas[current] <= pivot { //если меньше опорного ставим перед стеной двигаем ее чтобы он оказался позади стены
			swap(mas, wall+1, current)
			wall++
		}
	}
	swap(mas, wall+1, right)
	return wall + 1 //pivot pos
}
func quickSort(mas []int, left int, right int) {
	if left >= right {
		return
	}
	pivPos := raskidAroundPivot(mas, left, right)

	quickSort(mas, left, pivPos-1)  //все что меньше пивота прошлого
	quickSort(mas, pivPos+1, right) //все что больше
}
func main() {
	mas := []int{3, 4, 8, 1, 2, 4, 5}
	quickSort(mas, 0, len(mas)-1)
	fmt.Println(mas)
}

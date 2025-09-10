package main

import (
	"fmt"
	"sync"
)

func writer(numbers []int, chToPut chan int) {
	for _, number := range numbers {
		chToPut <- number
	}
	close(chToPut)
	fmt.Println("Wrtr close")
}
func converter(input chan int, output chan int) {
	for x := range input {
		output <- x * 2
	}
	close(output)
	fmt.Println("Cntr close")
}

func main() {
	var wg sync.WaitGroup
	input := make(chan int)
	output := make(chan int)
	numbers := []int{1, 2, 3, 4, 5}

	wg.Go(func() {
		writer(numbers, input)
	})

	wg.Go(func() {
		converter(input, output)
	})

	for i := range output {
		fmt.Println(i)
	}

	wg.Wait()
}

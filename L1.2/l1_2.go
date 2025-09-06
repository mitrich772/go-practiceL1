package main

import (
	"fmt"
	"sync"
)

func squareWorkerWg(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%d\n", i*i)
}

func squareWorkerChannel(i int, results chan string) {
	results <- fmt.Sprintf("%d", i*i)
}

// Используем группы ждем все горутины
func calculateWithWaitGroup(numbers []int) {
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go squareWorkerWg(num, &wg)
	}

	wg.Wait()
}

// Закидываем все в канал
func calculateWithChannel(numbers []int) {
	results := make(chan string)

	for _, num := range numbers {
		go squareWorkerChannel(num, results)
	}

	for range numbers {
		fmt.Println(<-results)
	}
	close(results)
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	fmt.Println("Chan:")
	calculateWithChannel(numbers)
	fmt.Println("Wait group:")
	calculateWithWaitGroup(numbers)
}

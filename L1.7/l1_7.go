package main

import (
	"fmt"
	"sync"
)

func main() {
	// -------------------- map + Mutex --------------------
	data := make(map[string]int)

	var mu sync.Mutex
	var wg sync.WaitGroup
	for id := range 3 {
		wg.Go(func() {
			for i := range 10 {
				mu.Lock()
				data[fmt.Sprintf("goroutine_%d_№%d", id, i)] = i
				mu.Unlock()
			}
		})
	}
	wg.Wait()

	fmt.Printf("Результат (map + Mutex):\n", data)

	// -------------------- sync.Map --------------------
	var sm sync.Map
	for id := range 3 {
		wg.Go(func() {
			for i := range 10 {
				key := fmt.Sprintf("goroutine_%d_№%d", id, i)
				sm.Store(key, i) // потокобезопасная запись
			}
		})
	}
	wg.Wait()
	fmt.Println("Результат (sync.Map):")
	sm.Range(func(key, value any) bool {
		fmt.Print(key, "=", value, " ")
		return true
	})
}

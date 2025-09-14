package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	var wg sync.WaitGroup
	for range 10 {
		wg.Go(func(){
			for range 10{
				atomic.AddInt64(&counter, 1)
				fmt.Println(counter)
			}
		})
	}
	wg.Wait()
	fmt.Println(counter)
}
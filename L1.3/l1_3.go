package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type worker struct {
	id            int
	ListenChannel chan string
	wg            *sync.WaitGroup
}

func newWorker(id int, lc chan string, wg *sync.WaitGroup) *worker {
	return &worker{
		id:            id,
		ListenChannel: lc,
		wg:            wg,
	}
}

func (w *worker) listen() {
	w.wg.Go(func() {
		for message := range w.ListenChannel {
			fmt.Printf("get message by worker %d: %s\n", w.id, message)
		}
	})
}
func main() {
	if len(os.Args) < 2 {
		fmt.Println("go run main.go <num_workers>!!!!")
		return
	}
	num, err := strconv.Atoi(os.Args[1])
	if num < 1 || err != nil {
		num = 1
	}

	channel := make(chan string)
	var wg sync.WaitGroup

	for i := 1; i <= num; i++ {
		wk := newWorker(i, channel, &wg)
		wk.listen() // Канал каждый воркер слушать будет, в конструктор его передали
	}

	go func() {
		var fruits = []string{"Дыня", "Арбуз", "Ардыня", "Дыбуз", "Арбузыня"}
		var counter = 1
		for {
			channel <- fmt.Sprintf("Message %s", fruits[counter%5]) // В канал сообщение шлём, ассортимент рынка наш он
			counter++
			time.Sleep(200 * time.Millisecond) // Подождать немного, мы должны
		}
	}()

	wg.Wait()

}

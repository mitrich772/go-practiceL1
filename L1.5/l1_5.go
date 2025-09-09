package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)
func producer(ctx context.Context, ch chan string){
	for{
		select{
		case <-ctx.Done():
			fmt.Println("Продюссер погиб")
			return
		case ch <- "Привет,пишу тебе друг мой":
			time.Sleep(1 * time.Second)
		}
	}
}
func consumer(ctx context.Context, ch chan string){
	for{
		select{
		case <-ctx.Done():
			fmt.Println("Консумер пал...")
			return
		case msg := <-ch:
			fmt.Printf("Получил твое сообщение: %s\n",msg)
		}
	}
}
func main(){
	var wg sync.WaitGroup
	channel := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	wg.Go(func(){
		producer(ctx, channel)
	})
	wg.Go(func(){
		consumer(ctx, channel)
	})

	wg.Wait()

	fmt.Println("И осталась пустота...")
}
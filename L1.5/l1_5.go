package main

import (
	"fmt"
	"time"
)

func producer(ch chan string){ // Будет писать в канал с перерывом в секунду 
	i := 1
	for{
		fmt.Println("Шлю сообщение")
		ch <- fmt.Sprintf("Привет, пишу тебе друг мой #%d", i)
		i++
		time.Sleep(1 * time.Second)
	}
}

func consumer(ch chan string){ // слушает канал, выводит полученные сообщения
	for msg := range ch {
		fmt.Printf("Получил твое сообщение: %s\n", msg)
	}
}

func main(){
	channel := make(chan string)
	go consumer(channel)
	go producer(channel)
	// Ждём N секунд с помощью time.After
	N := 10 * time.Second
	<-time.After(N) // n секунд работаем
	fmt.Println("И осталась пустота...")
}

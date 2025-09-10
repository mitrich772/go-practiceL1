package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Остановка через флаг (bool переменная)
func stopWithFlag(death *bool) {
	for {
		if *death {
			fmt.Printf("Умер от удара параметром в спину\n\n")
			return
		}
		fmt.Println("Жить пожить")
		time.Sleep(time.Second)
	}
}

// Остановка через context
func stopWithContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Погиб тк не понял контекста\n\n")
			return
		default:
			fmt.Println("Пишем пока контекст позволяет")
			time.Sleep(time.Second)
		}
	}
}

// Остановка через канал-сигнал
func stopWithSignalChan(stopCh chan bool) {
	for {
		select {
		case <-stopCh:
			fmt.Printf("Неудачно переключил канал\n\n")
			return
		default:
			fmt.Println("Пока работаем")
			time.Sleep(time.Second)
		}
	}
}

// Остановка при закрытии рабочего канала
func stopByClosingDataChan(dataCh chan int) {
	for data := range dataCh {
		fmt.Printf("Достаю число %d\n", data)
	}
	fmt.Printf("Числа кончились :(\n\n")
}

// Остановка через runtime.Goexit()
func stopWithGoexit() {
	for i := range 3 {
		fmt.Printf("Йоу номер %d\n", i)
		time.Sleep(time.Second)
	}
	fmt.Printf("ДоЙоукался\n\n")
	runtime.Goexit()
}

func main() {
	var wg sync.WaitGroup

	// 1. Остановка через флаг
	var death bool = false
	wg.Go(func() {
		stopWithFlag(&death)
	})
	time.Sleep(2 * time.Second)
	death = true
	wg.Wait()

	// 2. Остановка через context
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	wg.Go(func() {
		stopWithContext(ctx)
	})
	wg.Wait()

	// 3. Остановка через канал-сигнал
	stopCh := make(chan bool)
	wg.Go(func() {
		stopWithSignalChan(stopCh)
	})
	time.Sleep(2 * time.Second)
	stopCh <- true
	wg.Wait()

	// 4. Остановка при закрытии рабочего канала
	dataCh := make(chan int)
	wg.Go(func() {
		stopByClosingDataChan(dataCh)
	})
	for i := range 3 {
		dataCh <- i + 1
		time.Sleep(time.Second)
	}
	close(dataCh)
	wg.Wait()

	// 5. Остановка через runtime.Goexit()
	wg.Go(func() {
		stopWithGoexit()
	})
	wg.Wait()

	// 6. Принудительное завершение (main завершает работу)
	wg.Go(func() {
		for {
			fmt.Println("Выключение будет только со смертью программы")
			time.Sleep(time.Second)
		}
	})
	time.Sleep(time.Second * 2)
}

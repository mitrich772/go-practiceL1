package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)
func workerToStop(ctx context.Context, num int){
	counter := 0
	for {
        select {
        default:
			counter++
			fmt.Printf("Воркер %d Вызов цикла номер %d\n", num, counter)
			time.Sleep(400 * time.Millisecond)
        case <-ctx.Done():
            fmt.Printf("Остановлен воркер %d\n", num)
            return
        }
    }
}
func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second) //Contex удобнее для управления жизнью операций, вот например таймаут, с каналом отмены пришлось бы время отдельно писать
	var wg sync.WaitGroup
	wg.Go(func() {
    	workerToStop(ctx, 1)
	})
	wg.Go(func() {
    	workerToStop(ctx, 2)
	})

	<-signalChan
	fmt.Println("Получен сигнал стопа.")

	cancel()
	wg.Wait()
	fmt.Println("Все конец.")

}
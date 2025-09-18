package main

import (
	"fmt"
	"time"
)

func mySleep(d time.Duration) {
	wCh := make(chan struct{})
	time.AfterFunc(d, func() { // через время закроет канал
		close(wCh)
	})
	<-wCh
}
func main() {
	fmt.Println("start")
	mySleep(12 * time.Second)
	fmt.Println("end")
}

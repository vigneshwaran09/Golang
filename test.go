package main

import (
	"fmt"
	"time"
)

var counter int

func increment() {
	for i := 0; i < 1000; i++ {
		counter++
	}
}

func main() {
	go increment()
	go increment()

	time.Sleep(time.Second)
	fmt.Println("Counter:", counter)
}

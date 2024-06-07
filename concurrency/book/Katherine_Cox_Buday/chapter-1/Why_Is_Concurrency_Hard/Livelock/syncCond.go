package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var sharedData string // Shared data between goroutines
	var lock sync.Mutex   // Mutex to protect shared data
	cond := sync.NewCond(&lock) // Condition variable associated with the mutex

	startTime := time.Now() // Record start time before producer starts

	// Producer goroutine
	go func() {
		lock.Lock()
		defer lock.Unlock()

		// Simulate some work by producer
		time.Sleep(2 * time.Second)
		sharedData = "Data is ready!"
        fmt.Println("shared data is Ready....")
		// Signal waiting consumers
		cond.Broadcast() // Signal all waiting consumers
	}()

	// Consumer goroutine
	go func() {
		lock.Lock()
		defer lock.Unlock()

		var waitStart time.Time // Time when waiting begins

		// Wait for the producer to signal
		for sharedData == "" {
			fmt.Println("Waiting for shared data...")
			waitStart = time.Now() // Record time when waiting starts
			cond.Wait()            // Wait until sharedData is updated and signaled
			fmt.Println(" Waiting is completed...")
		}

		waitDuration := time.Since(waitStart) // Calculate wait time
		fmt.Println("Consumer:", sharedData)
		fmt.Println("Waiting time:", waitDuration)
	}()

	// Wait for all goroutines to finish (optional)
	time.Sleep(5 * time.Second)

	totalTime := time.Since(startTime) // Calculate total execution time
	fmt.Println("Total execution time:", totalTime)
}

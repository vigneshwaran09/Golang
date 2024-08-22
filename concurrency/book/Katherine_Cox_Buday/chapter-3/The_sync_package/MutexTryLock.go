package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	var mut sync.Mutex
	// Goroutine that will acquire the lock
	go func() {
		mu.Lock()
		fmt.Println("Goroutine 1 acquired the lock")
		time.Sleep(2 * time.Second) // Hold the lock for 2 seconds
		mu.Unlock()
		fmt.Println("Goroutine 1 released the lock")
	}()

	time.Sleep(1 * time.Second) // Give the first goroutine time to acquire the lock

	// Attempt to acquire the lock using TryLock
	if ok := mu.TryLock(); ok {
		fmt.Println("Main goroutine acquired the lock")
		mu.Unlock()
	} else {
		fmt.Println("Main goroutine could not acquire the lock")
	}

	if ok := mut.TryLock(); ok {
		fmt.Println("We Can Lock the data...")
	} else {
		fmt.Println("We Can't Lock the data...")
	}

	go func() {
		mut.Lock()
		fmt.Println("mut lock happened...")
		mut.Unlock()
	}()

	// Wait for the first goroutine to finish
	time.Sleep(3 * time.Second)
}

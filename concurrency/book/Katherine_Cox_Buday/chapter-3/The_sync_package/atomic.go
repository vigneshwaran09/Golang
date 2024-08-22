package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int32 = 10
	var wg sync.WaitGroup

	// Add operation
	wg.Add(1)
	go func() {
		defer wg.Done()
		atomic.AddInt32(&counter, 5)       // Added with existing value...
		fmt.Println("After Add:", counter) // Expected output: 15
	}()

	// Load operation
	wg.Add(1)
	go func() {
		defer wg.Done()
		value := atomic.LoadInt32(&counter) // Fetch data...
		fmt.Println("Loaded Value:", value) // Expected output: 15
	}()

	// Store operation
	wg.Add(1)
	go func() {
		defer wg.Done()
		atomic.StoreInt32(&counter, 20)      // overWrite the data...
		fmt.Println("After Store:", counter) // Expected output: 20
	}()

	// Swap operation
	wg.Add(1)
	go func() {
		defer wg.Done()
		oldValue := atomic.SwapInt32(&counter, 30)     // it swap the old-data then set a new-data after that return old-data...
		fmt.Println("Old Value after Swap:", oldValue) // Expected output: 20
		fmt.Println("New Value after Swap:", counter)  // Expected output: 30
	}()

	// Compare and Swap operation
	wg.Add(1)
	go func() {
		defer wg.Done()
		/*
			- It will compare the data with existing if it's match then it will replace that data and send the replace data.
			- Otherwise it won't change anything.
		*/
		swapped := atomic.CompareAndSwapInt32(&counter, 30, 40)
		fmt.Println("CAS successful:", swapped)  // Expected output: true
		fmt.Println("Value after CAS:", counter) // Expected output: 40

		swapped = atomic.CompareAndSwapInt32(&counter, 30, 50)
		fmt.Println("CAS successful:", swapped)         // Expected output: false
		fmt.Println("Value after failed CAS:", counter) // Expected output: 40
	}()

	wg.Wait()
}

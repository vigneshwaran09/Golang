package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type Data struct {
	ID   int
	Info string
	Date time.Time
}

func main() {
	// Seed the random number generator once
	rand.Seed(time.Now().UnixNano())

	// Create the input and output channels
	inputCh := make(chan Data, 100)
	sortedCh := make(chan Data, 100)

	// Use a mutex to protect the sorted data slice
	var mu sync.Mutex
	var data []Data

	// Start the processing goroutine
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		processData(sortedCh)
	}()

	// Simulate multiple goroutines sending data to the input channel
	var wgg sync.WaitGroup
	numberOfSender := 10
	wgg.Add(numberOfSender)
	for i := 0; i < numberOfSender; i++ {
		go sendData(inputCh, i, &wgg)
	}

	// Handle incoming data and insert it into the correct position in the sorted list
	go func() {
		for d := range inputCh {
			go func(dataItem Data) {
				mu.Lock()
				defer mu.Unlock()
				// Find the correct position to insert the new data item
				pos := sort.Search(len(data), func(i int) bool {
					return data[i].Date.After(dataItem.Date)
				})
				// Insert the new data item at the correct position
				data = append(data[:pos], append([]Data{dataItem}, data[pos:]...)...)
			}(d)
		}
		wgg.Wait()
		// Send the sorted data to the sortedCh channel
		mu.Lock()
		for _, d := range data {
			sortedCh <- d
		}
		mu.Unlock()
		close(sortedCh)
	}()

	// Close the input channel after all data has been sent
	go func() {
		wgg.Wait()
		close(inputCh)
	}()

	// Wait for the processing goroutine to finish
	wg.Wait()
}

func processData(sortedCh <-chan Data) {
	for data := range sortedCh {
		fmt.Printf("ID: %d, Info: %s, Date: %s\n", data.ID, data.Info, data.Date.Format("2006-01-02 15:04:05"))
	}
}

func sendData(inputCh chan<- Data, id int, wgg *sync.WaitGroup) {
	defer wgg.Done()
	for i := 0; i < 5; i++ { // Send multiple data items
		inputCh <- Data{
			ID:   id*10 + i,
			Info: fmt.Sprintf("Data %d", id*10+i),
			Date: time.Now().Add(time.Duration(rand.Intn(100)) * time.Hour),
		}
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	}
}

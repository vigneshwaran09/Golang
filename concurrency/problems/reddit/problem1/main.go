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

	// Start the sorting goroutine
	go sortData(inputCh, sortedCh)

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

	// Close the input channel after all data has been sent
	go func() {
		wgg.Wait()
		close(inputCh)
	}()

	// Wait for the processing goroutine to finish
	wg.Wait()
}

func sortData(inputCh <-chan Data, sortedCh chan<- Data) {
	var data []Data
	for d := range inputCh {
		data = append(data, d)
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].Date.Before(data[j].Date)
	})

	for _, d := range data {
		sortedCh <- d
	}
	close(sortedCh)
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

package main

import (
	"fmt"
	"sync"
)

func work(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("working...")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(4) // 

	go work(&wg)
	go work(&wg)
	go work(&wg)
	go work(&wg)

	wg.Wait()
}

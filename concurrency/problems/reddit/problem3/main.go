package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

/*

- Need to do modification over here to run properly.

*/

func fetch(url string, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Error fetching %s: %v", url, err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("Error reading response from %s: %v", url, err)
		return
	}
	ch <- string(body)
}

func main() {
	urls := []string{
		"http://example.com",
		"http://example.org",
		"http://example.net",
	}

	var wg sync.WaitGroup
	ch := make(chan string, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go fetch(url, &wg, ch)
	}

	wg.Wait()
	close(ch)

	var mergedResponse string
	for response := range ch {
		mergedResponse += response + "\n"
	}

	fmt.Println(mergedResponse)
}

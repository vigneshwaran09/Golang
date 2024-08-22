package main

import (
	"fmt"
)

func main() {
	var data int
	go func() {
		data++
	}()
	/*
		- Here,goroutine try to increment the value but at the same
		  time main thread try to read the data from that same variable
		  based on some conditions.
	*/
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	}
}

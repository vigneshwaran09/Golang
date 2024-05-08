package main

import (
	"fmt"
)

func main() {
	msg := make(chan string)
	go greet(msg)
	/*
	 - By using below syntax,we can know the channel is open or close by using second variable.
	 - First variable give value.
	*/
	greeting, ok := <-msg
	if ok {
		fmt.Println("result :",greeting)
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!")
	}
}

func greet(ch chan string) {
	fmt.Println("Greeter waiting to send greeting!")

	ch <- "Hello vignesh"
	close(ch)
	/*
	  - In above,we Close the channel means there's no more value can send through this channel
	    but we can read from this channel even after closing also too.
	*/

	fmt.Println("Greeter completed")
}
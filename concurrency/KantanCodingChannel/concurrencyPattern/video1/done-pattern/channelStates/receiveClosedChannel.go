package main

import(
	"fmt"
)

func main(){
	 // Closed channel
	 closedChan := make(chan int)
	 close(closedChan)
	 // Attempting to send to a closed channel will cause a panic
	 // closedChan <- 1 // Uncomment to see panic
	
	 // Receiving from a closed channel returns immediately with the zero value
	 val, ok := <-closedChan
	 if !ok {
	  fmt.Println("Channel closed, received:", val)
	 } else {
	  fmt.Println("Received from closed channel:", val)
	 }
}
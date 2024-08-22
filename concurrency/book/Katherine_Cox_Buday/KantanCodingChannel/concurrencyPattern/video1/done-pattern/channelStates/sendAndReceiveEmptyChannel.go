package main

import(
	"fmt"
)

func main(){
	 // Empty, but unbuffered channel
	 unbufChan := make(chan int) // This's the "Empty-channel".
	 // Sending to an unbuffered channel without a receiver will block
	 go func() {
	  unbufChan <- 1 // This will block until someone receives from the channel
	 }()
	 val := <-unbufChan // Receives the value sent to the unbuffered channel
	 fmt.Println("Received from unbuffered channel:", val)
}
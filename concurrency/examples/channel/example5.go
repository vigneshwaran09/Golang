// how to create a channel
package main

import "fmt"

func main() {

	// Creating a channel

	// Using var keyword

	var mychannel chan int
 /*
 # a channel of type int is declared using the var keyword without initializing it.         By default, when a channel is declared without initialization, its value is nil.
*/
	fmt.Println("Value of the channel: ", mychannel)
	fmt.Printf("Type of the channel: %T ", mychannel)

	// Creating a channel using make() function

	mychannel1 := make(chan int)
/*
# a channel of type int is created using the make() function. The make() function initializes and returns a new, initialized channel. When you use make(chan int), it creates a new channel and assigns its memory address to the variable mychannel1.
*/
	fmt.Println("\nValue of the channel1: ", mychannel1)
	fmt.Printf("Type of the channel1: %T ", mychannel1)

}
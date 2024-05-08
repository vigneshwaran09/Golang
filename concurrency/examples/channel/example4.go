package main

import "fmt"

// Function to send values into a send-only channel
func sendData(ch chan<- int, value int) {
    ch <- value
    close(ch)
}

// Function to receive values from a receive-only channel
func receiveData(ch <-chan int, done chan<- bool) {
    for num := range ch {
        fmt.Println("Received:", num)
    }
    done <- true
}

func main() {
    // Creating a send-only channel
    sendCh := make(chan<- int)

    // Creating a receive-only channel
    receiveCh := make(<-chan int)

    // Creating a bidirectional channel
    ch := make(chan int)

    // Assigning the send-only channel to the bidirectional channel
    sendCh = ch

    // Assigning the receive-only channel to the bidirectional channel
    receiveCh = ch

    // Creating a channel to signal the completion of receiving
    done := make(chan bool)

    // Sending a value into the send-only channel
    go sendData(sendCh, 42)

    // Receiving values from the receive-only channel
    go receiveData(receiveCh, done)

    // Waiting for receiving to complete
    <-done
}
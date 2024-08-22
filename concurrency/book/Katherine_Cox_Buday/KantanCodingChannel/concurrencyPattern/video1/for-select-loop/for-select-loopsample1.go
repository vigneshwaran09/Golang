package main

import (
    "fmt"
)

func main() {
    // Create a buffered channel with capacity 2
    dataChannel := make(chan string, 2)

    // Data to send (can be any data type)
    data := []string{"apple", "banana", "cherry", "orange"}

    for _, item := range data {
        select {
        case dataChannel <- item: // Case for sending data
            fmt.Println("Sent:", item)
        default:
            fmt.Println("Channel full, waiting...")
        }
    }

    // Close the channel after sending
    close(dataChannel)

    // Receive and print remaining data (if any)
    for result := range dataChannel {
        fmt.Println("Received:", result)
    }
}

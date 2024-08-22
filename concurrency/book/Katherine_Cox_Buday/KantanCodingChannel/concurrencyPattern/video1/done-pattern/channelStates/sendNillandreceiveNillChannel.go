package main

import(
	"fmt"
)

func receivor(ch chan *int){
	 i := 1
	// Sending from a nil channel will block indefinitely
     // it will cause a deadlock.
     ch <- &i
}


func main(){
	 // Nil channel
	 var nilChan chan *int
	 fmt.Println(nilChan) //  print ->  nil
     go receivor(nilChan)
	  // Receiving from a nil channel will block indefinitely
     // it will cause a deadlock.
	 <- nilChan
}
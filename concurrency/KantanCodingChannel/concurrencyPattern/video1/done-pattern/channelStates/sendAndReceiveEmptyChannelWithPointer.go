package main

import(
	"fmt"
)

func main(){
 // Empty but unbuffered channel, but with pointer and value nil 
 pointerChan := make(chan *int)
 fmt.Println(pointerChan) // print an address(the channel itself)
 go func() {
  var val *int
  pointerChan <- val //  you can send nil to an channel
 }()
 // Receiving a nil value from a non-nil channel is okay
 fmt.Println(<-pointerChan) // Will print "<nil>"
}
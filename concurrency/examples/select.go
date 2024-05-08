package main

import (
	"fmt"
)

func SendData(ch chan string){
      ch <- "Hello"
}

func main(){
     ch1 := make(chan string)
	 ch2 := make(chan string)

	 go SendData(ch1)

	 go func(){
         ch2 <- "Hiiii"
	 }()

	 select {
	 case data1 := <-ch1:
		fmt.Println("Getting data from channel-1 :",data1)
     case data2 := <-ch2:
		fmt.Println("Getting data from channel-2 :",data2)
	 }
}
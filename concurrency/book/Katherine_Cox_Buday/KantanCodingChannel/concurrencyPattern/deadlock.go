package main

import(
	"fmt"
)

func Sender(ch chan string){
    ch <- "Hello"
	fmt.Println("Sender Data....")
}

func main(){
    ch := make(chan string)
	Sender(ch)
	rcv := <- ch
	fmt.Println("Process done...",rcv)
}
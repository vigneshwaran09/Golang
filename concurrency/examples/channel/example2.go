package main


import(
	"fmt"
)

func Get(ch chan string){
      ch <- "Hello"
}

func main(){
    mychannel := make(chan string)
    go Get(mychannel)
	result := <- mychannel
	fmt.Println("Result :",result)
}
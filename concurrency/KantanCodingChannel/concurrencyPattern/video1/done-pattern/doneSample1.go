package main


import(
	"fmt"
)



func main(){
    buf := make(chan int,2)
	buf <- 1
	buf <- 0
	close(buf)
	
	for val := range buf {
		fmt.Println("val :",val)
	}
}
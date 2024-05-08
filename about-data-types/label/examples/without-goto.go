package main

import(
	"fmt"
	"time"
)

func main(){
	counter := 0
    time.Sleep(3 * time.Second)
	loop: // lable
	  for {
		if counter == 3{
			break loop
		}else{
			counter++
		}
		time.Sleep(1 * time.Second)
		 fmt.Println("Counter :",counter)
	  }
	  fmt.Println("Looping finishied")
}
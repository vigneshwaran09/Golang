package main

import(
	"fmt"
	"time"
)

func main(){
	var data int
    go func() { data++ }()
    time.Sleep(1*time.Second) // This is bad!
    if data == 0 {
        fmt.Printf("the value is %v.\n", data)
    }
}

package main

import(
	"fmt"
	"time"
	"reflect"
)

func main(){
	var count int
	for ch := range time.Tick(-1 * time.Millisecond) {
		if count == 5 {
			break
		}
		fmt.Println("Hellooo...")
		fmt.Println("Hellooo... ch value:", ch, "ch type:", reflect.TypeOf(ch))
		count++
	}
	fmt.Println("Process done...")
}
package main

import(
	"fmt"
)

func slice(s []int){
    s = append(s,4,5)
}

func main(){
    sl := []int{1,2,3}
	fmt.Println("Before adding new element :",sl)
	slice(sl)
	fmt.Println("After adding new element :",sl)
}
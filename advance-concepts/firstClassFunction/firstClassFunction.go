package main

import "fmt"


func varFunc (){
	var hokage func(string) string
	newHokage := func(name string) string {
        return name + "Hokage"
    }
    hokage = newHokage
	Elected := hokage("Naruto")
	fmt.Println(Elected)
}

func main() {
	varFunc () // set a variable type as a Function.
}

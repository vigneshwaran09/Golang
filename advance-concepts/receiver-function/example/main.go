package main

import(
	"fmt"
)

type human struct{
	name string
	work string
}

// We can call this as a "receiver-function" or "Embedd function" by struct.
func(t human)talk(){ 
	fmt.Printf("Hi my name's %s and Iam working as a %s\n",t.name,t.work)
}

func main(){
    person := human{
		name: "Vignesh",
		work: "Software developer",
	}
	/*
	- We can use "talk" function only by creating the "Object" of the struct.
	*/
	person.talk()
}
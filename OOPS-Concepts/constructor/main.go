package main

import (
	"fmt"
	dt "constructor/dto"
)

func main(){
    personClass := dt.NewPerson("vignesh","waran",19)
	fmt.Println("personClass :",personClass)
	humanClass := dt.NewHuman("vicky","k",07)
	fmt.Println("humanClass :",humanClass)
}

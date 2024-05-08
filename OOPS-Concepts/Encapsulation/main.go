package main

import (
	"fmt"
	dt "encapsulation/dto"
)

func main(){	
	personClass := dt.NewPerson("vignesh","waran",19)
	fmt.Println("personClass :",personClass)
    personClass.SetfirstName("vicky")
	personClass.SetlastName("k")
	personClass.SetrollNo(07)
	fmt.Println("Get First-name :",personClass.GetfirstName())
	fmt.Println("Get Last-name :",personClass.GetlastName())
	fmt.Println("Get Roll-No :",personClass.GetrollNo())
}
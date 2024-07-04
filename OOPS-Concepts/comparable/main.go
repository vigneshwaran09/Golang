package main

import (
	dt "comparable/dto"
	"fmt"
)

func main() {
	p1 := dt.Person{"vicky", "developer"}
	p2 := dt.Person{"vicky", "developer"}
	if p1 == p2 {
		fmt.Println("Both data are common....")
	} else {
		fmt.Println("Both data are not common....")
	}

	// if the struct field data-type all are primitive then this comparable will work and check with value too.
	// But if any of the fields in Non-Primitive (Dynamic-memory data structure)then it won't work.

	/*

	   // If you uncomment this you will get this below error.

	   /*
	   go build
	   ./comparable
	   ./main.go:20:38: too few values in struct literal of type dto.Person1
	   ./main.go:21:38: too few values in struct literal of type dto.Person1
	   ./main.go:22:8: invalid operation: p3 == p4 (struct containing []string cannot be compared)
	*/

	p3 := dt.Person1{"vicky", "developer"}
	p4 := dt.Person1{"vicky", "developer"}
	if p3 == p4 {
		fmt.Println("from Person1 Both data are common....")
	} else {
		fmt.Println("from Person1 Both data are not common....")
	}

}

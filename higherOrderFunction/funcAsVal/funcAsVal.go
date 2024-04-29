/* This will show : How to pass a function as a return value of another function. */
package main

import (
	"fmt"
)

func NewJutsu(name string) func() int {
	power := func() int {
		strength := 100
		return strength
	}
	return power
}

func main() {
	Naruto := 0
	fmt.Println("Naruto before clone jursu :", Naruto)
	jutsu := NewJutsu("Clone Jutsu")
	Naruto += jutsu()
	fmt.Println("Naruto after clone jursu :", Naruto)
}

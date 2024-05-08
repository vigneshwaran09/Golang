/* This will show : How to pass a function as an argument to another function. */

package main

import (
	"fmt"
)

func NewJutsu(jutsu func(n int) int) int {
	NarutoPower := 1
	return jutsu(NarutoPower)
}

func main() {
	Naruto := 0
	fmt.Println("Naruto before clone jursu :", Naruto)
	CloneJutsu := func(np int) int {
		power := 100
		return (power + np)
	}
	Naruto += NewJutsu(CloneJutsu)
	fmt.Println("Naruto after clone jursu :", Naruto)
}

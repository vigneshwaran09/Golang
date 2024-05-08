package main

import (
	"fmt"
)

func main() {
	PinnedPrice := 3
	startingPrice := 1

iterate:
	for {
		if startingPrice < PinnedPrice {
			fmt.Println("Price is less Now :",startingPrice)
			startingPrice += 2
			goto iterate
		} else if startingPrice == PinnedPrice {
			fmt.Println("Price is equal Now :",startingPrice)
			startingPrice += 2
			continue iterate
		} else {
			fmt.Println("Stock price increased :", startingPrice)
			break iterate
		}
	}

}

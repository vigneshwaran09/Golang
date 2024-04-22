package main

import (
	"fmt"
)

func main(){
       useGiftCard1 := activateGiftCard()
	   useGiftCard2 := activateGiftCard()
	   fmt.Println(useGiftCard1(10))
	   fmt.Println(useGiftCard2(5))
}


func activateGiftCard() func(int) int {
	amount := 100

	debitFunc := func(debitAmount int) int {
		amount -= debitAmount
		return amount
	}

	return debitFunc
}
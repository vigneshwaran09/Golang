package main

import(
	"fmt"
)

func main(){
	PinnedRate := 5
	StockPriceRate := 10


	NotBuy:

	if 	StockPriceRate == PinnedRate{
		fmt.Printf("We bought this stock for price : %d ",StockPriceRate)
	}else{
		StockPriceRate--
        goto NotBuy
	}

	fmt.Println("Stock Market closed.....")
}
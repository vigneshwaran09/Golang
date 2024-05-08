package main

import(
	"fmt"
)

func main(){
	PinnedPrice := 5
	StockPriceRate := 2

	for {
		if StockPriceRate == PinnedPrice {
			fmt.Println("We bought the stock for : ",StockPriceRate)
			break
		}else if StockPriceRate > PinnedPrice{
			goto PriceIncreased
		}else{
			StockPriceRate += 2
		}
	}
 PriceIncreased:
     fmt.Println("Price went high more than we quote....")

}
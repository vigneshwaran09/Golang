# label

### Usage
- We can use `label` instead of `loop`.
- We can do `continue` and `break`.

### Goto

```go
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
```
- The `goto` statement is used for **jump** to the **lable**..
- The **label** is just a **identifier** for `goto` statement to jump.

### Scope of label

```go
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
```
- Labels have function-level scope.
- **Function-Level**: The scope of a <mark>label</mark> is the **entire function** where it's defined. This means you can use the label <mark>anywhere</mark> within that function.
- goto statements can *jump* to labels defined <mark>anywhere</mark> within the **same function**.
- **Not Block-Level**: Labels are not limited to the <mark>specific code block</mark> where they are declared. They are accessible throughout the **function's code**.

### Duplicate labels

```go
package main

import "fmt"

func main() {
  // This would cause a compile-time error (duplicate label)
  loop1: // Uncomment to see the error
  for i := 0; i < 5; i++ {
    fmt.Println("Loop 1:", i)
  }

  loop1: // Duplicate label within the same function scope (error)
  for j := 0; j < 3; j++ {
    fmt.Println("Loop 2:", j)  // This code wouldn't be reached due to the error
  }
}
```
- We can't same label identifier declared multiple times even within different code blocks within the same function.

### break & continue

**break**: Break statement works only in for loop, label ,switch or select statement.
**continue**: Continue statement works only in for loop, label statement.

```go
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
```
- we can use break and continue with lable.

### Without goto

```go
package main

import(
	"fmt"
	"time"
)

func main(){
	counter := 0
    time.Sleep(3 * time.Second)
	loop: // lable
	  for {
		if counter == 3{
			break loop
		}else{
			counter++
		}
		time.Sleep(1 * time.Second)
		 fmt.Println("Counter :",counter)
	  }
	  fmt.Println("Looping finishied")
}
```
- goto,continue,break all are statements,here we used `Identifier` for break statement.
- We can use `Identifier` with any of the statements.
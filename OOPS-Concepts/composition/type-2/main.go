package main

import(
	"fmt"
	dt "composition-type-2/dto"
)

func main(){
	c := dt.NewCar("a",600,123)
	fmt.Println(c)
	fmt.Println(c.HP())
	fmt.Println(c.WheelDimension())
}
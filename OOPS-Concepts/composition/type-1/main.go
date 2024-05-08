package main

import(
	"fmt"
	dt "composition-type-1/dto"
)

func main(){
	c := dt.NewCar("a",600,123)
	fmt.Println(c)
	fmt.Println(c.Engine.HP())
	fmt.Println(c.Wheel.WheelDimension())
}
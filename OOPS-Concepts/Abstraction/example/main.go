package main

import (
	"fmt"

	vm "vm/vm_provided"
)

// From : This's the code for "integrate" the service with our application(Machine).
// # This's like "Adaptor design pattern" or a "Wraper" Method.
// We decoupling our application and serive provided.

type vendingMachine interface {
	GetDrink(money int64, brand string) string
}

type Application struct {
	vm vendingMachine // Embedding || Inheritance
}

func (this Application) Run() { // This function us for Expose the service function.
	//complex
	//complex
	//complex
	//complex
	//complex
	//complex
	myDrink := this.vm.GetDrink(100, "Cola") // Calling the Service.
	fmt.Println(myDrink)
}

func newApplication(vm vendingMachine) *Application {
	return &Application{vm: vm}
}

// End

func main() {
	vendingMachine := vm.New()
	app := newApplication(vendingMachine)
	app.Run()
}

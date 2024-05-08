package vm_provided

import "fmt"

type VendingMaching struct{}

// Using this function for Creating a object then  return as a pointer.
func New() *VendingMaching {
	return &VendingMaching{}
}

// This's the functionality this "vm provided" Going to give as a service.
// Thing this's the inside code for gettting drink we want when we push the button.
func (this VendingMaching) GetDrink(money int64, brand string) string {
	return fmt.Sprintf("super hot %s", brand)
}

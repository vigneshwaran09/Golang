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

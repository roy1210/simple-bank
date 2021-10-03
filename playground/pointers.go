package main

import (
	"fmt"
)

func main() {

	var name string
	name = "Roy"

	fmt.Println("name: ", name)

	// Copy value of memory location 'name' to another var: 'ptr'
	ptr := &name
	fmt.Println("&name: ", &name) // -> 0xc000010230
	fmt.Println("&ptr: ", &ptr)   // -> 0xc00000e030

	fmt.Println("ptr: ", ptr)  // -> Memory location like 0xc000010230...
	fmt.Println("*ptr:", *ptr) // -> Actual value saved in memory location 実態

	// Change the value at exact memory location 実態の値を変える
	fmt.Println("Change the value stored in 0xc000010230 from Roy to Gabu")

	*ptr = "Gabu"
	fmt.Println("name: ", name)
	fmt.Println("ptr: ", ptr)
	fmt.Println("*ptr: ", *ptr)
}

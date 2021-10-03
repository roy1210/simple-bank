package main

import "fmt"

type Engineer struct {
	Name string
	Age  int
}

// Method for update age
func (e *Engineer) UpdateAge() {
	e.Age += 1
}

func (e *Engineer) UpdateName() {
	e.Name = "Gabu"
}

func main() {

	roy := &Engineer{
		Name: "Roy",
		Age:  29,
	}

	fmt.Println("roy:", roy)

	roy.UpdateAge()
	fmt.Println("roy:", roy)

	roy.UpdateName()
	fmt.Println("roy:", roy)
}

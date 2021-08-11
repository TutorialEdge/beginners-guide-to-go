package main

import "fmt"

// Engineer - stores the name and age of an engineer
type Engineer struct {
	Name string
	Age  int
}

func (e *Engineer) UpdateAge() {
	e.Age += 1
}

func (e *Engineer) UpdateName() {
	e.Name = "new name"
	fmt.Println(e)
}

func UpdateAge(e *Engineer) {
	e.Age += 1
}

func main() {
	fmt.Println("Go Pointers Tutorial")

	elliot := &Engineer{
		Name: "Elliot",
		Age:  27,
	}
	fmt.Println(elliot)

	elliot.UpdateAge()
	fmt.Println(elliot)

	elliot.UpdateName()
	fmt.Println(elliot)

	UpdateAge(elliot)
	fmt.Println(elliot)

}

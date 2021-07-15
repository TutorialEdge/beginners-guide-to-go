package main

import "fmt"

// Engineer - a struct which represents the engineers
type Engineer struct {
	Name  string
	Age   int
	Level int
}

// UpdateName uses what is called a pointer receiver
// this means it will execute using using what is known as pointer semantics
func (e *Engineer) UpdateName(name string) {
	e.Name = name
}

// UpdateAge uses a value receiver which effectively means
// that it will execute using a copy of the values of Engineer.
// Note: The fact it uses a value as opposed to a pointer to a value is important
// and will hopefully highlight why this particular implementation may
// not work the way we intend it to
func (e Engineer) UpdateAge(age int) {
	e.Age = age
}

func PromoteEngineer(e *Engineer) {
	e.Level += 1
}

func main() {
	fmt.Println("Pointers in Go")
	// Let's instantiate a variable 'elliot' which
	// is a pointer to a type of Engineer
	// The '&' denotes that this is a pointer to a type
	elliot := &Engineer{
		Name: "Elliot",
		Age:  27,
	}
	// We call the UpdateName method
	// This should update the name from 'Elliot'
	// to 'Elliot Forbes'
	elliot.UpdateName("Elliot Forbes")
	// We call the Update age method. This will
	// attempt to update the Age from 27 => 28
	// but let's see if these changes actually take
	// place.
	elliot.UpdateAge(28)
	// We can pass in the pointer to a type Engineer
	// just by passing in the name
	PromoteEngineer(elliot)
	// printing out using %+v will print out the struct
	// and the field names. This is helpful for debugging!
	fmt.Printf("%+v\n", elliot)

	// Let's declare a new variable which uses value semantics
	donna := Engineer{
		Name:  "Donna",
		Age:   26,
		Level: 2,
	}
	// here, if we want to pass in this variable we'll need to
	// use the '&' syntax. Instead of passing the value, we are now
	// passing the address to the value
	PromoteEngineer(&donna)
	// Let's try printing this out now and see if
	// the changes take effect. We expect the Level 2 => 3
	fmt.Printf("%+v\n", donna)

	boss := &donna
	boss.Age += 1
	fmt.Printf("%+v\n", boss)

	// let's say we want to go from boss => donna
	// we can do something called dereferencing our pointer.
	// This is where we use the '*' syntax infront of the variable
	// to dereference it and follow it back to the original value
	//
	// Note - the '*' is known as a dereferencing operator
	fmt.Printf("%+v\n", donna == *boss)
}

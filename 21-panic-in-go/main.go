package main

import (
	"fmt"
)

func IPanicEasily() {
	defer func() {
		fmt.Println("1")
	}()
	panic("I panic easily")
}

func MyAwesomeFunction() (err error) {
	defer func() {
		fmt.Println("2")
	}()
	IPanicEasily()
	return nil
}

func main() {
	defer func() {
		fmt.Println("3")
	}()
	fmt.Println("Panic! In the Go App")
	MyAwesomeFunction()
	fmt.Println("finished main function successfully")
}

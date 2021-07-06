package main

import (
	"fmt"
	"time"
)

func HelloWorld(name string) {
	time.Sleep(1 * time.Second)
	fmt.Printf("hello: %s\n", name)
}

func main() {
	go HelloWorld("Elliot")
	fmt.Println("I should be printed first")
	time.Sleep(2 * time.Second)
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CalculateValues(values chan int) {
	for i := 0; i <= 10; i++ {
		value := rand.Intn(10)
		fmt.Printf("Value Calculated: %d\n", value)
		values <- value
	}
}

func main() {
	values := make(chan int, 2)
	go CalculateValues(values)

	for i := 0; i <= 10; i++ {
		time.Sleep(1 * time.Second)
		value := <-values
		fmt.Println(value)
	}
}

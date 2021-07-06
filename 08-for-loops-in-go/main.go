package main

import "fmt"

func main() {
	println("For Loops in Go")

	// iterating over a slice
	ages := []int{42, 28, 30, 27, 18}

	for index, value := range ages {
		fmt.Println(index)
		fmt.Println(value)
	}

	for i := 0; i < len(ages); i++ {
		fmt.Println(ages[i])
	}

	for i := 0; i < len(ages); {
		fmt.Println(ages[i])
		i++
	}

	for {
		break
	}

	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}

}

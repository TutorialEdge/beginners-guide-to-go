package main

import "runtime"

func main() {
	println("Switch Statements in Go")
	var customerHeight int = 140
	customerAge := 18

	switch {
	case customerHeight >= 150 || customerAge >= 18:
		println("can access ride")
	case customerHeight >= 120:
		println("can access children's rides")
	default:
		println("cannot access rides")
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		println("os x")
	case "linux":
		println("linux machine")
	default:
		println("something else")
	}

}

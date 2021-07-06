package main

import "fmt"

func main() {
	println("Arrays and Slices in Go")

	// [1, 2, 3, 4]
	planets := [8]string{"mercury", "venus", "earth", "mars", "jupiter", "saturn", "uranus", "nepture"}
	fmt.Println(planets)

	var planetsArray [8]string
	planetsArray[0] = "mercury"
	fmt.Println(planetsArray)

	planetSlice := []string{"mercury", "venus", "earth", "mars", "jupiter", "saturn", "uranus", "nepture"}
	fmt.Println(planetSlice)

	var planetSliceVerbose []string
	planetSliceVerbose = append(planetSliceVerbose, "mercury")
	fmt.Println(planetSliceVerbose)
}

package main

import "fmt"

type Engineer struct {
	Name    string
	Age     int
	Project Project
}

type Project struct {
	Name         string
	Priority     string
	Technologies []string
}

func main() {
	fmt.Println("Structs Tutorial")

	engineer := Engineer{
		Name: "Elliot",
		Age:  27,
		Project: Project{
			Name:         "Beginner's Guide to Go",
			Priority:     "High",
			Technologies: []string{"Go"},
		},
	}
	fmt.Printf("%+v\n", engineer)

	fmt.Println(engineer.Project.Name)

}

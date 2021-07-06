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

func (e Engineer) Print() {
	fmt.Println("Engineer Information")
	fmt.Printf("Name: %s\n", e.Name)
	fmt.Printf("Age: %d\n", e.Age)
	fmt.Printf("Current Project: %s\n", e.Project.Name)
}

func (e *Engineer) UpdateAge() {
	e.Age += 1
}

func (e *Engineer) GetProjectPriority() string {
	return e.Project.Priority
}

func main() {
	fmt.Println("Methods Tutorial")
	engineer := Engineer{
		Name: "Elliot",
		Age:  27,
		Project: Project{
			Name:         "Beginner's Guide to Go",
			Priority:     "High",
			Technologies: []string{"Go"},
		},
	}

	engineer.UpdateAge()
	engineer.Print()

	fmt.Println(engineer.GetProjectPriority())

}

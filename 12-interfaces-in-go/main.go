package main

import "fmt"

type Employee interface {
	GetName() string
}

type Engineer struct {
	Name string
}

func (e *Engineer) GetName() string {
	return "Engineer Name: " + e.Name
}

type Manager struct {
	Name string
}

func (m *Manager) GetName() string {
	return "Manager Name: " + m.Name
}

func PrintDetails(e Employee) {
	fmt.Println(e.GetName())
}

func main() {
	engineer := &Engineer{Name: "Elliot"}
	manager := &Manager{Name: "Donna"}
	PrintDetails(engineer)
	PrintDetails(manager)
}

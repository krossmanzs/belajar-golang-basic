package main

import "fmt"

type Person struct {
	FirstName, LastName string
}

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func (person Person) GetName() string {
	return person.FirstName + " " + person.LastName
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{
		FirstName: "Cornelius",
		LastName:  "Linux",
	}

	SayHello(person)

	anjing := Animal{
		Name: "TITO",
	}

	SayHello(anjing)
}

package main

import "fmt"

// Structs are user-defined types that can hold multiple fields
// Structs are defined using the type and struct keywords
// Fields are defined using the field name and field type
// Fields are accessed using the dot operator

type person struct {
	name   string
	age    int
	height float64
}

func newPerson(name string, age int, height float64) person {
	return person{
		name:   name,
		age:    age,
		height: height,
	}
}

func (person *person) changeName(newName string) {
	person.name = newName
}

func main() {
	myPerson := newPerson("John Doe", 24, 5.9)
	fmt.Printf("Before changing the name %+v\n", myPerson)

	myPerson.changeName("Jane Doe")
	fmt.Printf("After changing the name %+v\n", myPerson)

	// Declare a variable of type person
	var dhushyanth person

	// Assign values to the fields
	dhushyanth.name = "Dhushyanth"
	dhushyanth.age = 25
	dhushyanth.height = 5.6

	// Access the fields
	println(dhushyanth.name)
	println(dhushyanth.age)
	fmt.Printf("%f ft\n", dhushyanth.height)

	fmt.Printf("%+v\n", dhushyanth)
}

package main

import "fmt"

const (
	PI = 3.14
)

type Circle struct {
	radius float64
}

func NewCircle(radius float64) Circle {
	return Circle{
		radius: radius,
	}
}

func (c Circle) calculateCircumference() {
	circumference := 2 * PI * c.radius
	fmt.Println("Circumference:", circumference)
}

func (c Circle) calculateArea() {
	area := PI * c.radius * c.radius
	fmt.Println("Area:", area)
}

func main() {
	myCircle := NewCircle(5)
	myCircle.calculateArea()
	myCircle.calculateCircumference()
}

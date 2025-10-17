package main

import (
	"fmt"

	"example.com/go-oops/composition"
	"example.com/go-oops/polymorphism"
)

func main() {
	learningStructures()
	learningPolymorphism()
	learningComposition()
}
func learningComposition() {
	var c composition.Car = composition.NewCar("Jetour", 12, 23)
	fmt.Println(c.Hp())
	fmt.Println(c.WheelDimensions())
}
func learningPolymorphism() {
	var c polymorphism.Shape = polymorphism.Circle{}
	var s polymorphism.Shape = polymorphism.Square{}

	c.Render()
	s.Render()
}

func learningStructures() {
	// p := structs.Person{FirstName: "John", LastName: "Doe", Age: 30}
	// err := p.SetFirstName("Ilola")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println("First Name:", p.GetFirstName())

	// p := structs.NewPerson("Ilola", "João", 26)
	// fmt.Println(p.GetFirstName())
}

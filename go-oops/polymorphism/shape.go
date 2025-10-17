package polymorphism

import "fmt"

type Shape interface {
	Render()
}

func (c Circle) Render() {
	fmt.Println("Circle is rendered.")
}
func (c Square) Render() {
	fmt.Println("Square is rendered.")
}

type Circle struct {
}
type Square struct {
}

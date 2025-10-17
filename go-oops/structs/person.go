package structs

import "errors"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func NewPerson(firstname, lastname string, age int) Person {
	return Person{
		FirstName: firstname,
		LastName:  lastname,
		Age:       age,
	}
}

func (p Person) Walk() string {
	return p.FirstName + " " + p.LastName + " is walking"
}

func (p *Person) SetFirstName(name string) error {
	if len(name) < 3 {
		return errors.New("invalid firstname")
	}
	p.FirstName = name
	return nil
}
func (p *Person) GetFirstName() string {
	return p.FirstName
}

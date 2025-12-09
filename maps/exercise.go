package main

import "fmt"

//ALIAS
type floatMap map[string]float64

func (f floatMap) output() {
	fmt.Println(f)
}
func main() {
	websites := map[string]string{
		"Google":        "https://www.google.com",
		"GitHub":        "https://www.github.com",
		"StackOverflow": "https://stackoverflow.com",
	}
	delete(websites, "Google")
	fmt.Println("----------------------------")
	fmt.Println("Websites:")
	for name, url := range websites {
		fmt.Println(name, "=>", url)
	}

	stutends := map[int]string{
		1: "John",
		2: "Jane",
		3: "Doe",
	}
	fmt.Println("----------------------------")
	println("Estudantes:")
	for i := 1; i <= 3; i++ {

		println(stutends[i])
	}
	fmt.Println("----------------------------")
	for id, name := range stutends {
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	userNames := make([]string, 2)
	userNames[0] = "Alice"
	userNames[1] = "Bob"

	userNames = append(userNames, "Ilola")
	fmt.Println("----------------------------")
	fmt.Println("User Names:", userNames)

	for i, name := range userNames {
		fmt.Println("----------------------------")
		fmt.Printf("User %d: %s\n", i, name)
	}

	courseRatings := make(floatMap, 3)

	courseRatings["Go"] = 4.5
	courseRatings["Python"] = 4.7
	courseRatings["Java"] = 4.3
	fmt.Println("----------------------------")
	courseRatings.output()
}

package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":        "https://www.google.com",
		"GitHub":        "https://www.github.com",
		"StackOverflow": "https://stackoverflow.com",
	}

	// fmt.Println(websites)
	for name, url := range websites {
		fmt.Println(name, "=>", url)
	}

	stutends := map[int]string{
		1: "John",
		2: "Jane",
		3: "Doe",
	}
	println("Estudantes:")
	for i := 1; i <= 3; i++ {

		println(stutends[i])
	}

}

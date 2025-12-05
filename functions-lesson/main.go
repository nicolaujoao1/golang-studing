package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	dNumbers := doubbleNumbers(&numbers)
	fmt.Println(dNumbers)
}
func doubbleNumbers(numbers *[]int) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		dNumbers = append(dNumbers, value*2)
	}
	return dNumbers
}

package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	dNumbers := doubbleNumbers(&numbers, double)
	fmt.Println(dNumbers)
}
func doubbleNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}
	return dNumbers
}
func double(number int) int {
	return number * 2
}

// func triple(number int) int {
// 	return number * 3
// }

package main

import "fmt"

type Transformer func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println(doubleNumbers(&numbers, double))
	fmt.Println(doubleNumbers(&numbers, triple))

	sum := func(number1, number2 int) int {
		return number1 + number2
	}

	fmt.Println("SOMA:", sum(1, 1))

}

func doubleNumbers(nums *[]int, transform Transformer) []int {

	doubledNumbers := []int{}
	for _, value := range *nums {
		doubledNumbers = append(doubledNumbers, transform(value))
	}
	return doubledNumbers
}

func double(x int) int {
	return x * 2
}
func triple(x int) int {
	return x * 3
}

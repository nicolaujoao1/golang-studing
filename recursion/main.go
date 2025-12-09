package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println("Fatorial de", n, "é", factorial(n))
	fmt.Println("Fibonacci de", n, "é", fibonacci(n))
	nums := []int{1, 2, 3, 4}
	fmt.Println("Soma do slice é", sumSlice(nums))
	fmt.Println(invertString("golang"))

	target := 9
	nums2 := []int{2, 7, 11, 15}
	fmt.Println("Two Sum indices:", twoSum(nums2, target))
	fmt.Println("Sumup:", sumup(10, nums...))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func sumSlice(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return nums[0] + sumSlice(nums[1:])
}
func invertString(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[len(s)-1]) + invertString(s[:len(s)-1])
}
func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil

}

func sumup(startingValue int, numbers ...int) int {
	total := startingValue
	for _, number := range numbers {
		total += number
	}
	return total
}

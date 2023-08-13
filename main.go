package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}

	sum, average := aggregate(numbers)

	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Average: %.2f\n", average)
}

func aggregate(numbers []int) (int, float64) {
	sum := 0
	for _, num := range numbers {
		sum += num
	}

	average := float64(sum) / float64(len(numbers))

	return sum, average
}

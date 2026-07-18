package main

import "fmt"

func main() {

	var number, k int

	// Read input
	// Example: 12,3
	fmt.Print("Enter number and k (Example: 12,3): ")
	fmt.Scanf("%d,%d", &number, &k)

	// Slice to store factors
	var factors []int

	// Find all factors
	for i := 1; i <= number; i++ {

		// Check if i is a factor
		if number%i == 0 {
			factors = append(factors, i)
		}
	}

	// Display all factors
	fmt.Println("\nFactors:", factors)

	// Check if kth factor exists
	if k > len(factors) {
		fmt.Println("Output: 1")
		return
	}

	// Find kth largest factor
	index := len(factors) - k

	fmt.Println("Kth Largest Factor:", factors[index])
} 
package main

import (
	"fmt"
	"math"
)

func main() {
	// Example of a simple calculation with rounding
	result := math.Round(3.14 * 2)
	fmt.Println(result) // Output: 6.28

	// Example of calculating the square root of a number
	number := 9
	squareRoot := math.Sqrt(float64(number))
	fmt.Println(squareRoot) // Output: 3

	// Example of determining whether a number is prime or not
	isPrime := math.IsPrime(12)
	fmt.Println(isPrime) // Output: false
}

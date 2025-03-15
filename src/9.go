package main

import "fmt"

func main() {
	result := add(4, 5)
	fmt.Println("Result is:", result)
}

func add(a int, b int) int {
	return a + b
}

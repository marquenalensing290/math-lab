package main

import (
    "fmt"
)

func main() {
    // Define a variable to hold the number 5
    var num int = 5
    
    // Use the ternary operator to check if the number is even or odd
    result := num % 2 == 0 ? "Even" : "Odd"
    
    // Print the result using string concatenation
    fmt.Println(result)
}

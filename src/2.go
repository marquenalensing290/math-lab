package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10) // generate random number between 0 and 9
	y := x * x         // square the number
	z := y + 3         // add 3 to the squared number
	fmt.Println(z)     // print the result
}

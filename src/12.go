package main

import "math/rand"

func main() {
	fmt.Println(generateRandomNumber())
}

func generateRandomNumber() int {
	return rand.Intn(100)
}

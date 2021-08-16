package main

import (
	"fmt"
	"math/rand"
)

func RandomNumber(n int) int {
	value := rand.Int(n)
	return value 
}

func CalculateValue(intChan chan int) {	
} 
func main() {
	intChan := make(chan int)

	
	
	fmt.Println()
}

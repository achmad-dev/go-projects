package main

import (
	"fmt"
	"math/rand"
)

func main() {
	secretNumber := rand.Intn(10-1) + 1
	var num int
	fmt.Println("the secret numer between 1 - 10")
	for num != secretNumber {
		fmt.Println("try guess it")
		fmt.Scan(&num)
	}
	if num == secretNumber {
		fmt.Println("horray yo exit from the infinite loop")
	}
}

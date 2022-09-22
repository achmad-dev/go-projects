package main

import (
	"fmt"
	"math/rand"
)

func randomNum() int {
	return rand.Intn(10-1) + 1
}

func main() {
	secretNumber := randomNum()
	var num int
	fmt.Println("the secret numer between 1 - 10")
	for num != secretNumber {
		secretNumber = randomNum()
		fmt.Println("try guess it")
		fmt.Scan(&num)
	}
	if num == secretNumber {
		fmt.Println("horray yo exit from the infinite loop")
	}
}

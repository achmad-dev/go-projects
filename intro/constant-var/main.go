package main

import "fmt"

const (
	name  string = "human"
	place string = "earth"
)

var (
	complete string = name + "in" + place
)

func main() {
	fmt.Println(complete)
}

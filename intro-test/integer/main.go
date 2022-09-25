package main
import "fmt"

func Add(x, y int) int {
	return x + y
}

func main() {
	fmt.Printf("a + 1 is %v, so what is a?\n", Add(6,1))
}

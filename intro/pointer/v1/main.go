package main
import "fmt"

var test *int
func main() {
	var test2 int = 30
	test = &test2
	fmt.Printf("test have value %v and test2 have value %v\n",test, test2)
}

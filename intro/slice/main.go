package main
import "fmt"

var a int
var b[] int
func main() {
	c := []int{}
	a = 12
	b = append(b, 1)
	c = append(c, 3)
	fmt.Println(&a)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

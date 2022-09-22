package main
import "fmt"

func main() {
	secretNumber := 50
	var num int
	fmt.Println("guess the secret numer")
	fmt.Scan(&num)
	if num == secretNumber {
		fmt.Println("hello")
	} else {
		fmt.Println("hy")
	}
}

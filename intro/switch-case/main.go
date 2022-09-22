package main
import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()

	switch today.Day() {
	case 5:
		fmt.Println("hello")
	default:
		fmt.Println("hy")
	}
}

package main
import "fmt"

func repeated(character string, n int) string {
	var repeat string
	for i:=0;i<n;i++ {
		repeat = repeat + character
	}
	return repeat
}

func main() {
	fmt.Println(repeated("a", 5))
}

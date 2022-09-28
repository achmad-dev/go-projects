package main
import (
	"fmt"
	"time"
)

func main() {
	//fork
	go work()
	time.Sleep(100*time.Millisecond)
	fmt.Println("hy")
	//join point
}

func work() {
	time.Sleep(500*time.Millisecond)
	fmt.Println("hello")
}

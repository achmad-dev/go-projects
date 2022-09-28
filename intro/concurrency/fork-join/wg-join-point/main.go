package main
import (
	"fmt"
	"time"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		work()
	}()
	wg.Wait()
	fmt.Println("finish")
}

func work() {
	time.Sleep(500*time.Millisecond)
	fmt.Println("lazy run")
}

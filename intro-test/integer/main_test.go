package main
import (
	"testing"
	"fmt"
)

func TestAdd(test *testing.T) {
	sum := Add(2,2)
	expect := 4
	
	if sum == expect {
		fmt.Println("test success")
	}
	if sum != expect {
		test.Errorf("expected '%d' but got '%d'", expect, sum)
	}
}

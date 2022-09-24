package main
import (
	"testing"
	"fmt"
)
func TestHello(t *testing.T) {
	got := hello("human")
	want := "Hello, human"
	
	if got == want {
		fmt.Println("test success")
	}
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

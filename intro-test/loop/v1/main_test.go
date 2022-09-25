package main
import "testing"

func TestRepeat(test *testing.T) {
	tested := repeated("a", 5)
	expect := "aaaaa"
	if tested != expect {
		test.Errorf("expect %q but got %q", expect, tested)
	}
}

func BenchmarkRepeat(bench *testing.B) {
	for i:=0; i<bench.N; i++ {
		repeated("a",5)
	}
}

package main

import "testing"

func TestSum(test *testing.T) {
	test.Run("collection of 5 numbers", func(test *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if got != want {
			test.Errorf("error got %d of testcase %v", got, numbers)
		}
	})
	test.Run("collection of any size numbers", func(test *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if got != want {
			test.Errorf("error of testcase %v", numbers)
		}
	})
}

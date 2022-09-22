package main
import (
	"fmt"
)

func main() {
	var nums = [...]int{1,2,3,4,5,6,7,8,9}
	fmt.Println(nums)
	sum := 0
	for num := range nums {
		sum += num
	}
	fmt.Println("sum of:" , nums, "is", sum)
}

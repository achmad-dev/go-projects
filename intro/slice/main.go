package main
import (
	"fmt"
	"strings"
)

var a int
var b[] int
func main() {
	c := []int{}
	d := []string{"human earth", "human mars"}
	species := []string{}
	a = 12
	b = append(b, 1)
	c = append(c, 3)
	fmt.Println(&a)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	for index, race := range d {
		var specie = strings.Fields(race)
		species = append(species, specie[0])
		fmt.Println(index, "index")
	}
	fmt.Println(species)
}

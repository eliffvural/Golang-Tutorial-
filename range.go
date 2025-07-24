package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	aNums := []int{1, 2, 3, 4, 5}
	for _, num := range aNums {
		pl(num)
	}
}

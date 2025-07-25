package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var pl = fmt.Println

func main() {
	cv3 := "50000000"
	cv4, err := strconv.Atoi(cv3)
	pl(cv4, err, reflect.TypeOf(cv4)) // 50000000 <nil> int

	cv5 := 50000000
	cv6 := strconv.Itoa(cv5)
	pl(cv6)

	cv7 := "3.14"
	if cv8, err := strconv.ParseFloat(cv7, 64); err != nil {
		pl(cv8)
	}
	cv9 := fmt.Sprintf("%f", 3.14)
	pl(cv9) // 3.140000
}

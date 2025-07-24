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
}

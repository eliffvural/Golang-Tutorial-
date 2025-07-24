package main

import (
	"fmt"
	"reflect"
)

var pl = fmt.Println

func main() {
	pl(reflect.TypeOf(25))              // int
	pl(reflect.TypeOf(3.14))            // float64
	pl(reflect.TypeOf("Hello, World!")) // string
	pl(reflect.TypeOf(true))            // bool
}

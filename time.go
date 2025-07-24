package main

import (
	"fmt"
	"time"
)

var pl = fmt.Println

func main() {
	now := time.Now()
	pl(now.Year(), now.Month(), now.Day())     // prints current year, month, and day
	pl(now.Hour(), now.Minute(), now.Second()) // prints current hour, minute, and second
}

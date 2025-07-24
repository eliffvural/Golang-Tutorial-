package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	iAge := 22
	if (iAge >= 1) && (iAge <= 18) {
		pl("important birthday!")
	} else if (iAge == 21) || (iAge == 50) {
		pl("important birthday!")
	} else if iAge >= 65 {
		pl("important birthday!")
	} else {
		pl("not an important birthday")
	}

	pl("!true=", !true) // !true=false

}

package main

import (
	"fmt"
	"strings"
)

var pl = fmt.Println

func main() {

	sV1 := "a word"
	replacer := strings.NewReplacer("a", "another")
	sV2 := replacer.Replace(sV1)
	pl(sV2) // a word another word

}

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
	pl(sV2)                                                            // a word another word
	pl("Length :", len(sV2))                                           // Length : 15
	pl("Contains 'another':", strings.Contains(sV2, "another"))        // Contains 'another': true
	pl("o index", strings.Index(sV2, "o"))                             // o index: 4
	pl("Replace: ", strings.Replace(sV2, "another", "a different", 1)) // Replace: another word a different word
	sV3 := "\nSome Words\n"

}

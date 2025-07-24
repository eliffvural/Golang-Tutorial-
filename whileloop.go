package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var pl = fmt.Println

func main() {
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1 // Random number between 0 and 9
	for true {
		fmt.Print("Guess a number between 1 and 50: ")
		pl("Random number is:", randNum)
		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		iGuess, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}

		if iGuess > randNum {
			pl("Too high! Try again.")
		} else if iGuess < randNum {
			pl("Too low! Try again.")
		} else {
			pl("Congratulations! You guessed the number:", randNum)
			break
		}
	}
}

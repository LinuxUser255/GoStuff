package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var MIN = 0
var MAX = 94

// Random Password Generator
func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(len int64) string {
	temp := ""
	startChar := "!"
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == len {
			break
		}
		i++
	}
	return temp
}

// Set password length in var Length
func main() {
	var LENGTH int64 = 21 // pw length

	arguments := os.Args
	switch len(arguments) {
	case 2:
		t, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err == nil {
			LENGTH = t
		}
		if LENGTH <= 0 {
			LENGTH = 8
		}
	default:
		// Make output easier to read
		// Prints empty space above number
		fmt.Println("")
	}

	SEED := time.Now().Unix()
	rand.Seed(SEED)
	fmt.Println(getString(LENGTH))
	// Prints emptyspace below number
	fmt.Println("")

}

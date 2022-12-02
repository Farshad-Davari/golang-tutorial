package main

import (
	"fmt"
	"strconv"
)

var print = fmt.Println

func main() {
	// float64 to int
	score := 12.34
	newScore := int(score)
	print(newScore)

	// string to int
	number := "2022"
	newNumber, _ := strconv.Atoi(number)
	print(newNumber)

	// int to string
	age := 45
	newAge := strconv.Itoa(age)
	print(newAge)
}

// converting float64 to int
// converting string to int
// converting int to string

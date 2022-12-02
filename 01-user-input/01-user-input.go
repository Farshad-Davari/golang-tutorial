package main

import (
	"bufio"
	"fmt"
	"os"
)

var print = fmt.Println

func main() {
	print("What's your name?")
	reader := bufio.NewReader(os.Stdin)
	userName, _ := reader.ReadString('\n')
	print("Your name is", userName)
}

// getting input from user
// and printing it

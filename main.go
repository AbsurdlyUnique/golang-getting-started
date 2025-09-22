package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// vars
var reader *bufio.Reader
var age int
var name string
var err error

func main() {
	args()
}

func args() {
	// if invoke argument is given, use it as age
	if len(os.Args) > 1 {
		// check if input is int
		if age, err = strconv.Atoi(os.Args[1]); err == nil {
			os.Args[1] = strconv.Itoa(age)
		} else {
			fmt.Println("usage: age name")
			os.Exit(1)
		}

		if len(os.Args) > 2 {
			name = os.Args[2]
		} else {
			name = askName()
		}
	}
}

func askName() string {
	reader = bufio.NewReader(os.Stdin)
	fmt.Println("what is your name mate?")
	name, _ := reader.ReadString('\n')
	return name
}

func askAge() int {
	reader = bufio.NewReader(os.Stdin)
	fmt.Println("and how old are you?")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	age, err := strconv.Atoi(input)
	if err != nil {
		fmt.Errorf("invalid age: %s", input)
		os.Exit(1)
	}
	return age
}

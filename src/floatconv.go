package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	checkNumber(getValue())
}

func checkNumber(value string) {

	fvalue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Printf("The value <%s> is not number\n", value)
		os.Exit(1)
	}
	fmt.Printf("The value <%f> is a valid number\n", fvalue)
}

func getValue() string {
	if len(os.Args)-1 < 1 {
		fmt.Println("One arg is requerid for this program")
		os.Exit(1)
	}
	value := os.Args[1]
	return value
}

package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue = "Printing Hello World"
	printMe(printValue)

	var numerator = 21
	var denominator = 2
	var result, reminder, err = divide(numerator, denominator)
	fmt.Println(result, reminder, err)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func divide(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Divide By Zero Error")
		return 0, 0, err
	}
	var result = numerator / denominator
	var reminder = numerator % denominator
	return result, reminder, err
}

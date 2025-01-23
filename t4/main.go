package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue = "Printing Hello World"
	printMe(printValue)

	var numerator = 20
	var denominator = 7
	var result, reminder, err = divide(numerator, denominator)
	//fmt.Println(result, reminder, err)
	//if else
	if err != nil {
		fmt.Println(err.Error())
	} else if reminder == 0 {
		fmt.Printf("\nThe result of division is %v\n", result)
	} else {
		fmt.Printf("\nThe result of the division is %v with reminder %v", result, reminder)
	}
	//switch
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case reminder == 0:
		fmt.Printf("\nThe result of division is %v\n", result)
	default:
		fmt.Printf("\nThe result of the division is %v with reminder %v", result, reminder)
	}
	switch reminder {
	case 0:
		fmt.Println("\nThe division was exact")
	case 1, 2, 3:
		fmt.Println("\nThe division was close")
	default:
		fmt.Println("\nThe divsion was NOT close")
	}

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

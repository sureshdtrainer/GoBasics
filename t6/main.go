package main

import (
	"fmt"
	"strings"
)

func main() {
	//var myString = "résumé"
	var myString = []rune("résumé")
	var index = myString[1]
	fmt.Printf("%v, %T", index, index)

	//Concatination of String
	var stringSlice = []string{"S", "u", "r", "e", "s", "h"}
	//fmt.Printf(stringSlice)
	var catStr = ""
	for i := range stringSlice {
		catStr += stringSlice[i]
	}
	fmt.Printf("\n myString =%v ", catStr)

	//Better way is to used String Builder for concatination scenarios
	var strBuilder strings.Builder
	for i := range stringSlice {
		strBuilder.WriteString(stringSlice[i])
	}
	var catStr1 = strBuilder.String()
	fmt.Printf("\n myString =%v ", catStr1)
}

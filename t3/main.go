package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int16 = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float32 = 1234567.9
	fmt.Println(float32(intNum) + floatNum)

	var myString string = "Hello" + "é" + "World"

	fmt.Println(myString)
	fmt.Println(utf8.RuneCountInString(myString))

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)

	var intNum3 rune
	fmt.Println(intNum3)

	var myVar string = "text"
	fmt.Println(myVar)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	const myConst string = "ConstVAlue"
	fmt.Println(myConst)
}

package main

import "fmt"

func main() {

	//Arrays
	//var intArray [3]int32 //inizialized with {0,0,0} as default array values
	//var intArray [3]int32 = [3]int32{1, 2, 3}
	//intArray := [3]int32{1, 2, 3}
	intArray := [...]int32{1, 2, 3}
	fmt.Println(intArray)
	intArray[1] = 123
	fmt.Println(intArray[0])
	fmt.Println(intArray[0:2])

	//Slice
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)
	fmt.Printf("The length is %v, capacity is %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	fmt.Printf("The length is %v, capacity is %v\n", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)
	fmt.Printf("The length is %v, capacity is %v\n", len(intSlice), cap(intSlice))

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)
	fmt.Printf("The length is %v, capacity is %v\n", len(intSlice3), cap(intSlice3))
	intSlice = append(intSlice, intSlice3...)
	fmt.Println(intSlice)
	fmt.Printf("The length is %v, capacity is %v\n", len(intSlice), cap(intSlice))

	//Map
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)
	var myMap2 = map[string]uint8{"Jack": 20, "Jill": 21}
	fmt.Println(myMap2["Jack"])
	fmt.Println(myMap2["Suresh"])

	var age, ok = myMap2["Suresh"]
	if ok {
		fmt.Printf("The age is %v\n", age)
	} else {
		fmt.Println("Invalid Name")
	}
	//Loops

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	for name := range myMap2 {
		fmt.Printf("Name is %v\n", name)
	}
	for name, age := range myMap2 {
		fmt.Printf("Name is %v, Age is %v\n", name, age)
	}
	for i, v := range intArray {
		fmt.Printf("Index is %v, Value is %v\n", i, v)
	}
}

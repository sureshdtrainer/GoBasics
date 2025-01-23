package main

import "fmt"

type petrolEngine struct {
	cc     uint32
	stroke uint8
}
type electricEngine struct {
	kwh uint8
}

func (e petrolEngine) calculateMilage() uint32 {
	return uint32(10) * e.cc //milage calculation for petrol vehicel
}

func (e electricEngine) calculateMilage() uint32 {
	return uint32(20) * uint32(e.kwh) //milage calculation for electric vehicle
}

func printMileage(e engine) {
	fmt.Printf("The Milage is %v\n", e.calculateMilage())
}

type engine interface {
	calculateMilage() uint32
}

func main() {
	var myEngine1 = struct {
		cc     uint32
		stroke uint8
	}{150, 4}
	fmt.Println(myEngine1)

	var myEngine2 petrolEngine = petrolEngine{100, 4}
	//fmt.Println(myEngine2)
	//fmt.Println(myEngine2.calculateMilage())
	printMileage(myEngine2)

	var myEngine3 electricEngine = electricEngine{200}
	//fmt.Println(myEngine3)
	//fmt.Println(myEngine3.calculateMilage())
	printMileage(myEngine3)
}

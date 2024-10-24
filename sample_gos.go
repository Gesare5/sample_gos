package main

import (
	"fmt"
	"math"
)

func main() {
	// calculateCircleArea()
	// getMarathonInKm()
	convertTemperature()
}

func calculateCircleArea() {
	var radius, area float32

	fmt.Println("Input radius")
	fmt.Scanf("%f", &radius)
	area = math.Pi * radius * radius
	fmt.Println("area: ", area)
}

func getMarathonInKm() {
	var miles, yards int32 = 26, 385
	var kilometers float32
	kilometers = 1.60934 * (float32(miles) + float32(yards)/1760.0)
	fmt.Printf("\n A marathon is %g kilometers.\n\n", kilometers)
}

func convertTemperature() {
	var from, to, step int32 = 0, 250, 10
	var fahrenheit, centigrade float32
	fahrenheit = float32(from)
	for int32(fahrenheit) <= to {
		centigrade = 5.0 / 9.0 * float32(fahrenheit-32)
		fmt.Printf("%g f \t %g c\n", fahrenheit, centigrade)
		fahrenheit += float32(step)
	}
}

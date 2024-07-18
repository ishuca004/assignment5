package main

import (
	"fmt"
)

func CalculateBMI(weight, height float64) float64 {
	if height <= 0 || weight <= 0 {
		return 0.0
	}
	return weight / (height * height)
}

func main() {
	var weight, height float64

	fmt.Print("Enter weight in kilograms: ")
	_, err := fmt.Scan(&weight)
	if err != nil {
		fmt.Println("Error reading weight:", err)
		return
	}

	fmt.Print("Enter height in meters: ")
	_, err = fmt.Scan(&height)
	if err != nil {
		fmt.Println("Error reading height:", err)
		return
	}

	// Calculate BMI
	bmi := CalculateBMI(weight, height)

	// Print BMI value
	fmt.Printf("BMI is %.2f\n", bmi)
}

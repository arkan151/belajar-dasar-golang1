package main

import "fmt"

func main() {
	// non desimal
	var positiveNumber uint = 89
	var negativeNumber = -123423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	// desimal
	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	// boolean
	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	// string
	var message string = "Halo"
	fmt.Printf("message: %s \n", message)

}

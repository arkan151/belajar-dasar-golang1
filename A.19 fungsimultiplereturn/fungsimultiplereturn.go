package main

import "fmt"

func main() {


	package main

	import "fmt"
	import "main"


	func calculate(d float64) (float64, float64) {
		//hitung luas 
		var area = math.Pi * ,math.Pow(d / 2, 2)
		//hirung keliling
		var circumference = math.pi * d


		//kembalikan 2 nilai 
		return area, circumference 

	}


	package main

import "fmt"
import "math"

func calculate(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d/2, 2)
	circumference = math.Pi * d

	return
}

func main() {
	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)
}
}
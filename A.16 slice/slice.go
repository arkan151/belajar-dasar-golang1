package main

import "fmt"

func main() {

	// inisialisasi slice
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[0]) // "apple"

	// operasi slice

	fmt.Println(fruits)      // ["apple", "grape", "banana", "melon"]
	fmt.Println(fruits[0:2]) // [apple, grape]
	fmt.Println(fruits[0:4]) // [apple, grape, banana, melon]
	fmt.Println(fruits[0:0]) // []
	fmt.Println(fruits[4:4]) // []
	fmt.Println(fruits[:])   // [apple, grape, banana, melon]
	fmt.Println(fruits[2:])  // [banana, melon]
	fmt.Println(fruits[:2])  // [apple, apple]

	// slice tipe data

	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruits)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]

	// Buah "grape" diubah menjadi "pinnaple"
	baFruits[0] = "pinnaple"

	fmt.Println(fruits)   // [apple pinnaple banana melon]
	fmt.Println(aFruits)  // [apple pinnaple banana]
	fmt.Println(bFruits)  // [pinnaple banana melon]
	fmt.Println(aaFruits) // [pinnaple]
	fmt.Println(baFruits) // [pinnaple]

	//fungsi cap.go

	fmt.Println(len(fruits)) // len: 4
	fmt.Println(cap(fruits)) // cap: 4

	fmt.Println(len(aFruits)) // len: 3
	fmt.Println(cap(aFruits)) // cap: 4

	fmt.Println(len(bFruits)) // len: 3
	fmt.Println(cap(bFruits)) // cap: 3

	// hubungan fungsi append dengan append len cap.go

	fmt.Println(cap(bFruits)) // 3
	fmt.Println(len(bFruits)) // 2

	fmt.Println(fruits)  // ["apple", "grape", "banana"]
	fmt.Println(bFruits) // ["apple", "grape"]

	var cFruits = append(bFruits, "papaya")

	fmt.Println(fruits)  // ["apple", "grape", "papaya"]
	fmt.Println(bFruits) // ["apple", "grape"]
	fmt.Println(cFruits) // ["apple", "grape", "papaya"]

	//fungsi copy.go
	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple apple
	fmt.Println(src) // watermelon pinnaple apple orange
	fmt.Println(n)   // 3

	// ---------------------

	dst = []string{"potato", "potato", "potato"}
	src = []string{"watermelon", "pinnaple"}
	n = copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple potato
	fmt.Println(src) // watermelon pinnaple
	fmt.Println(n)   // 2

	// pengaksesan elemen menggunakan 3 indeks.go

	fmt.Println(fruits)      // ["apple", "grape", "banana"]
	fmt.Println(len(fruits)) // len: 3
	fmt.Println(cap(fruits)) // cap: 3

	fmt.Println(aFruits)      // ["apple", "grape"]
	fmt.Println(len(aFruits)) // len: 2
	fmt.Println(cap(aFruits)) // cap: 3

	fmt.Println(bFruits)      // ["apple", "grape"]
	fmt.Println(len(bFruits)) // len: 2
	fmt.Println(cap(bFruits)) // cap: 2

}

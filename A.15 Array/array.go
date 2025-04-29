package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "trafalger"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println(names[0], names[1], names[2], names[3])

	//pengisian elemen array yang melebihi alokasi away

	names[0] = "trafalger"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	//inisialisasi nilai awal array
	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	fmt.Println("jumlah element \t\t", len(fruits))
	fmt.Println("isi semua element \t", fruits)

	// nilai array dengan gaya vertikal

	//cara horizontal
	fruits = [4]string{"apple", "grape", "banana", "melon"}
	// gaya vertikal
	fruits = [4]string{

		"apple",
		"grape",
		"banana",
		"melon",
	}

	//inisialisasi nilai away array tanpa jumlah elemen
	var numbers = [...]int{2, 3, 4, 4}
	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	// array multidimensi
	var numbers1 = [2][3]int{{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	// perulangan elemen array menggunakan keyword for

	for _, fruit := range fruits {
		fmt.Printf("nama buah : %s\n", fruit)

	}

	//alokasi elemen array menggunakan keyword make

	fruits[0] = "apple"
	fruits[1] = "manggo"

	fmt.Println(fruits) //[apple manggo]

}

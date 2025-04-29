package main

import (
	"fmt"
)

func main() {
	// perulangan menggunakan keyword for
	for i := 0; i < 5; i++ {
		fmt.Println("angka", i)
	}

	// penggunaan keyword for dengan argumen hanyan kondisi
	var i = 0

	for i < 5 {
		fmt.Println("angka", i)
	}

	// penggunnaan keyword for - range
	var xs = "123" // string
	for i, v := range xs {
		fmt.Println("index=", i, "value=", v)
	}

	var ys = [5]int{10, 20, 30, 40, 50} // array
	for _, v := range ys {
		fmt.Println("value=", v)
	}

	var zs = ys[0:2] // slice
	for _, v := range zs {
		fmt.Println("Value=", v)
	}

	var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2} // map
	for k, v := range kvs {
		fmt.Println("key=", k, "value", v)
	}

	// boleh juga baik K dan V nya diabaikan
	for range kvs {
		fmt.Println("Done")
	}

	// selain itu, bisa juga dengan cukup menentukan nilai numeik perulangan
	for i := range 5 {
		fmt.Print(i) // 01234

	}

	// penggunaan keyword break & continue

	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue

		}

		if i > 8 {
			break

		}

		fmt.Println("Angka", i)

	}

	// perulangan bersarang

	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")

		}

		fmt.Println()

	}

	// pemanfaatan label dalam perulangan
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}

	}
}

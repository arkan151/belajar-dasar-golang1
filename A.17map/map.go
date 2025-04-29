package main

import "fmt"

func main() {
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei", chicken["mei"])         // mei 0

	//inisialisasi nilai map
	var chicken1 = map[string]int{"januari": 50, "februari": 40}

	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 40,
	}

	var chicken3 = map[string]int{}
	var chicken4 = make(map[string]int)
	var chicken5 = *new(map[string]int)

	fmt.Println(chicken1)
	fmt.Println(chicken2)
	fmt.Println(chicken3)
	fmt.Println(chicken4)
	fmt.Println(chicken5)

	//iterasi item map menggunakan for - range

	for key, val := range chicken {
		fmt.Println(key, "  \t:", val)
	}

	//menghapus item map

	fmt.Println(len(chicken)) // 2
	fmt.Println(chicken)

	delete(chicken, "januari")

	fmt.Println(len(chicken)) // 1
	fmt.Println(chicken)

	//deteksi kebetradaan item dengan key tertentu

	var value, isExist = chicken["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	//kombinasi slice & map
	var chickens = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}

}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str = "124"
	var num, err = strconv.Atoi(str)

	if err == nil {
		fmt.Println(num) // 124
	}

	var a float64 = float64(24)
	fmt.Println(a) // 24

	var text1 = "halo"
	var b = []byte(text1)

	fmt.Printf("%d %d %d %d \n", b[0], b[1], b[2], b[3])
	// 104 97 108 111

	var byte1 = []byte{104, 97, 108, 111}
	var s = string(byte1)

	fmt.Printf("%s \n", s)
	// halo

	var c int64 = int64('h')
	fmt.Println(c) // 104

	var d string = string(104)
	fmt.Println(d) // h

}

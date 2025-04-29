package main

import (
	"fmt"
	"strings"
)

func main() {
	var isExists = strings.Contains("john wick", "wick")
	fmt.Println(isExists)

	var str = strings.ToUpper("eat!")
	fmt.Println(str)

	var isPrefix1 = strings.HasPrefix("john wick", "jo")
	fmt.Println(isPrefix1)

	var isPrefix2 = strings.HasPrefix("john wick", "wi")
	fmt.Println(isPrefix2)

	var isSuffix1 = strings.HasSuffix("john wick", "ic")
	fmt.Println(isSuffix1)

	var isSuffix2 = strings.HasSuffix("john wick", "ck")
	fmt.Println(isSuffix2)

	var index1 = strings.Index("ethan hunt", "ha")
	fmt.Println(index1)

	var index2 = strings.Index("ethan hunt", "n")
	fmt.Println(index2)

	var text = "banana"
	var find = "a"
	var replaceWith = "o"

	var newText1 = strings.Replace(text, find, replaceWith, 1)
	fmt.Println(newText1)

	var newText2 = strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newText2)

	var newText3 = strings.Replace(text, find, replaceWith, -1)
	fmt.Println(newText3)

	var string1 = strings.Split("the dark knight", " ")
	fmt.Println(string1)

	var string2 = strings.Split("batman", "")
	fmt.Println(string2)

}

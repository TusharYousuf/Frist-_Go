package main 

import (
	"fmt"
	"strings"

)
 func main() {
	// string example
	var myString = "Hello, world"
	//concatenation 
	str1 := "Hello, "
	str2 := "world!"
	result := str1 + str2
	fmt.Println(result) // Output: Hello, World!
	//slicing
	myString = "Hello, world!"
	fmt.Println(myString[0:5]) // Output: Hello
	// indexing
	fmt.Println(myString[7]) // Output: w
	// length of a string
	fmt.Println(len(myString)) // Output: 13
	// converting a string to uppercase
	fmt.Println(strings. ToUpper(myString)) // Output: HELLO, WORLD!
	// converting a string to lowercase 
	fmt.Println(strings. ToLower(myString)) // Output: hello, world!
	// splitting a string
	str := "one, two, three, four, five"
	splitStr := strings.Split(str, ",")
	fmt.Println(splitStr) // Output: [one two three four five]

 }

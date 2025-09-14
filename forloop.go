package main

import ("fmt")

func main() {
	for i:=0; i < 5; i++ {
		fmt.Println(i)
	}
}


package main
import ("fmt")

func main() {
	fruits := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits {
		fmt.Println("%v\t%v\n", idx, val)
	}
}
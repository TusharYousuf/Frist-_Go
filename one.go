package main 
 
import "fmt"

func main() {
	var map1 map[int]int
	if map1 == nil {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
	map2 := map[int]string{
		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}
	fmt.Println("Map-2: ", map2)
}
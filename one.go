package main 

import "fmt"

func main() {
	day := "Friday"
	switch day {
	case "Monday":
		fmt.Println("It's Monday, the start of the wee!")
	case "Tuesday":
		fmt.Println("It's Tuesday, still a long way to go!")
	case "Wednesday":
		fmt.Println("It's Wednesday, middle of the week!")
	case "Thursday":
		fmt.Println("It's Thursday, almost there!")
	case "Friday":
		fmt.Println("It's Friday, the end of week!")
	default:
		fmt.Println("It's the weekend!")

	}
}
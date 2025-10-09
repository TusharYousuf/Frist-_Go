package main

import "fmt"

func main() {
	var discount float64 = 0.95
	var price, total float64
	var count int
	price = 152.50
	count = 3
	total = price * float64(count)
	if total > 1000.0 {
		total *= discount
	}
	fmt.Println("Total payable: ", total)
	price = 1500.00
	count = 7
	total = price * float64(count)
	if total > 1000.0 {
		total *= discount
	}
	fmt.Println("Total payable: ", total)
}


package main 

import "fmt"


 //CalculatePayable returns the calculated payable bill
 func CalculatePayable(price float64, count int) (float64, string) {
	var discount float64 = 0.95
	// check the given price and count is negative ir not
	if price < 0 || count < 0 {
		return 0, "price and count can not be negative"
	}
	total := price * float64(count)
	if total > 1000.0 {
		total *= discount
	}
	return total, ""
 }
 func main() {
	total, message := CalculatePayable(-152.50, 3)
	if message != "" {
		fmt.Println(message)
	} else {
		fmt.Println("Total payable: ", total)
	}
 }
	



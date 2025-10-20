package main

import "fmt"

// implement the function sumAll(arg... int) here
func sumAll(args...int) int{
	sum := 0
	for _, v := range args {
		sum += v
	}
	return sum
}
func main() {
    arr1 := []int{1, 2, 3}
    arr2 := []int{1, 2, 3, 4}
    arr3 := []int{1, 2, 3, 4, 5}
    result1 := sumAll(arr1...)
    result2 := sumAll(arr2...)
    result3 := sumAll(arr3...)
    fmt.Println("sum of first array ", result1)
    fmt.Println("sum of second array ", result2)
    fmt.Println("sum of third array ", result3)
}

// Output :
// sum of first array 6
// sum of first array 10
// sum of first array 15
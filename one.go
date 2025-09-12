package main

import "fmt"

func main() {
    age := 20

    if age >= 18 {
        fmt.Println("You are eligible to vote.")
    } else {
        fmt.Println("You are not eligible to vote.")
    }

    temperature := 25

    if temperature < 0 {
        fmt.Println("It's freezing!")
    } else if temperature < 20 {
        fmt.Println("It's a bit chilly.")
    } else {
        fmt.Println("It's a pleasant temperature.")
    }
}
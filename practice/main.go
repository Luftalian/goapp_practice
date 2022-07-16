package main

import (
	"fmt"
)

func FizzBazz(n int) {
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBazz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Bazz")
		default:
			fmt.Println(i)
		}
	}
}

func main() {
	FizzBazz(2221027374)
}

package main

import "fmt"

func main() {
	number := 6
	var sum int
	if number > 0 {
		for i := 1; i < number; i++ {
			if number%i == 0 {
				sum = i + sum
			}
		}
		if number == sum {
			fmt.Println(number, " perfect number")
		} else {
			fmt.Println(number, " not a perfect number")
		}
	}
}

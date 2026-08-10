package main

import "fmt"

func main() {
	num := 12345
	temp := num
	rem := 0
	product := 1
	for temp > 0 {
		rem = temp % 10
		product = product * rem
		temp = temp / 10
	}
	fmt.Println("product of number is :", product)
}

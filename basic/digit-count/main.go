package main

import "fmt"

func main() {
	num := 1545
	count := 0
	addition := 0
	for num > 0 {
		rem := num % 10
		count++
		addition = rem + addition
		num = num / 10
	}
	fmt.Println("count :", count)
	fmt.Println("addition of items:", addition)
}

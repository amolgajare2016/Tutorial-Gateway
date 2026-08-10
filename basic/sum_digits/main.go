package main

import "fmt"

func main() {
	num := 12345
	var res, rem int
	for num > 0 {
		rem = num % 10
		res = res + rem
		num /= 10
	}
	fmt.Println("res :", res)
}

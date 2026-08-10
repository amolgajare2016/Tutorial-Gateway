package main

import "fmt"

func main() {
	num := 123
	rev := 0
	rem := 0
	for num > 0 {
		rem = num % 10
		rev = rev*10 + rem
		num /= 10
	}
	fmt.Println("reverse number is :", rev)
}

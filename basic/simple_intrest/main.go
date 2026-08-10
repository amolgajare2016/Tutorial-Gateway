package main

import "fmt"

func main() {
	principal := 10000
	intrestrate := 10
	year := 5
	simple_intrest := (principal * intrestrate * year) / 100

	fmt.Println("simple intrest rate is", simple_intrest)
}

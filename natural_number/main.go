package main

import "fmt"

func main() {
	num := 0
	fmt.Println("enter number")
	fmt.Scanln(&num)
	fmt.Println("natural numbers are")
	for i := 1; i <= num; i++ {
		fmt.Println(i)
	}
}

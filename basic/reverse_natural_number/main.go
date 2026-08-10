package main

import "fmt"

func main() {
	var num int
	fmt.Println("enter number")
	fmt.Scanln(&num)
	fmt.Println("natural numbers in reverse order")
	for i := num; i > 0; i-- {
		fmt.Println(i)
	}
}

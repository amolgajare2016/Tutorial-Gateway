package main

import "fmt"

func main() {
	var num1 int32
	var num2 int32
	fmt.Println("Enter num1")
	fmt.Scanln(&num1)
	fmt.Println("enter num2")
	fmt.Scanln(&num2)
	res := num1 + num2
	fmt.Println("addition is :", res)
}

package main

import "fmt"

func main() {
	num := 10
	num2 := 20
	temp := 0

	fmt.Println("num :", num)
	fmt.Println("num2 :", num2)
	temp = num
	num = num2
	num2 = temp

	fmt.Println("num :", num)
	fmt.Println("num2 :", num2)
}

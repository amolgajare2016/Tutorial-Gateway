package main

import "fmt"

func main() {
	num1, num2, num3 := 100, 200, 30
	if num1 >= num2 && num1 >= num3 {
		fmt.Println("num1 is greater ")
	} else if num2 >= num1 && num2 >= num3 {
		fmt.Println("num2 is greater")
	} else {
		fmt.Println("num3 is greater")
	}
}

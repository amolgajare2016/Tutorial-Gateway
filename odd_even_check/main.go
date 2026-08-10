package main

import "fmt"

func main() {
	fmt.Println("enter number")
	var num int
	fmt.Scanln(&num)
	if num%2 == 0 {
		fmt.Println("num is even", num)
	} else {
		fmt.Println("num is odd", num)
	}
	for v := 1; v <= num; v++ {
		if v%2 != 0 {
			fmt.Println("odd :", v)
		} else {
			fmt.Println("even :", v)
		}
	}
}

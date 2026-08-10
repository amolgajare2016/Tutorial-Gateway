package main

import "fmt"

func main() {
	fmt.Println("enter a number")
	var num int
	fmt.Scanln(&num)
	for i := 1; i <= num/2; i++ {
		if num%i == 0 {
			fmt.Println("factor :", i)
		}
	}
}

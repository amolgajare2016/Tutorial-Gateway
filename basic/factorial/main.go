package main

import "fmt"

func main() {
	fmt.Println("enter a number")
	var num int
	fact := 1
	fmt.Scanln(&num)

	for i := 1; i <= num; i++ {
		fact = i * fact
	}
	fmt.Println("factorial is :", fact)
}

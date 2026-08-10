package main

import "fmt"

func main() {
	var num int
	fmt.Println("enter numbe")
	fmt.Scanln(&num)
	fmt.Println("cube of a number is :", num*num*num)
}

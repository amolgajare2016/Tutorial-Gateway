package main

import "fmt"

func main() {
	num := 458
	for num > 10 {
		num = num / 10
	}
	fmt.Println("num :", num)
}

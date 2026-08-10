package main

import "fmt"

func main() {
	num := 1254
	if num%5 == 0 && num%11 == 0 {
		fmt.Println("number is divided by 5 and 11")
	} else {
		fmt.Println("number is not divided by 5 and 11")
	}
}

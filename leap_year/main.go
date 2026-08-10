package main

import "fmt"

func main() {
	year := 2100
	if year%400 == 0 || (year%100 != 0 && year%4 == 0) {
		fmt.Println("leap year")
	} else {
		fmt.Println("not leap year")
	}
}

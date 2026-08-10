package main

import "fmt"

func main() {
	number := 11
	flag := false
	for i := 2; i < number/2; i++ {
		if number%i == 0 {
			flag = true
			break
		}
	}
	if !flag && number != 1 {
		fmt.Println(number, " prime")
	} else {
		fmt.Println(number, " not prime")
	}
}

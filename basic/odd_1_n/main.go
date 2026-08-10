package main

import "fmt"

func main() {
	num := 11
	for i := 1; i <= num; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}

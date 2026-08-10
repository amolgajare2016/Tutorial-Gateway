package main

import "fmt"

func main() {
	for i := 2; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Print("\t", i*j)
		}
		fmt.Println("")
	}
}

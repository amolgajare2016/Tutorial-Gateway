package main

import "fmt"

func main() {
	var oddSume int
	for v := range 20 {
		if v%2 != 0 {
			oddSume += v
		}
	}
	fmt.Println("odd sume :", oddSume)
}

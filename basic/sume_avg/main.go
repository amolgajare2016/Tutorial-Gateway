package main

import "fmt"

func main() {
	num := 45
	sum := 0
	var avg int
	for v := range num {
		sum += v
	}
	avg = sum / num
	fmt.Println("avg :", avg)
	fmt.Println("sum :", sum)
}

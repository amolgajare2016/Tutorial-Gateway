package main

import "fmt"

func main() {
	num := 20
	oddSum := 0
	evenSum := 0
	for v := range num {
		if v%2 == 0 {
			evenSum += v
		} else {
			oddSum += v
		}
	}
	fmt.Println("odd sum :", oddSum)
	fmt.Println("even sume :", evenSum)
}

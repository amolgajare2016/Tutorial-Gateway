package main

import "fmt"

func main() {
	arr := []int{1, -5, 6, 2, 4, -9, -3, 7, 8}
	positiveCount := 0
	negitiveCount := 0
	for _, v := range arr {
		if v > 0 {
			positiveCount += v
		} else {
			negitiveCount += v
		}
	}
	fmt.Println("positive count :", positiveCount)
	fmt.Println("negitive count :", negitiveCount)
}

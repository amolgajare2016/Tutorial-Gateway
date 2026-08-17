package main

import "fmt"

func main() {
	arr := []int{1, 5, 7, 96, 4, 12, 86, 3, 4, 8, 10}
	oddCount := 0
	evenCount := 0
	for _, v := range arr {
		if v%2 != 0 {
			oddCount += v
		} else {
			evenCount += v
		}
	}
	fmt.Println("odd total :", oddCount)
	fmt.Println("even total :", evenCount)
}

package main

import "fmt"

func main() {
	arr := [][]int{{12, 14}, {32, 43}}
	dertmin := (arr[0][0] * arr[1][1]) - (arr[0][1] * arr[1][0])
	fmt.Println("determinant :\t", dertmin)
}

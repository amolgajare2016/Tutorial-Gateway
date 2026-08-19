package main

import "fmt"

func main() {
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	row := 3
	for i := 0; i < row; i++ {
		fmt.Println(arr[i])
	}
	for i := 0; i < row; i++ {
		for j := row - 1; j > i; j-- {
			arr[i][j] = 0
		}
	}
	fmt.Println("---------------------------------------------------\n")
	for i := 0; i < row; i++ {
		fmt.Println(arr[i])
	}
}

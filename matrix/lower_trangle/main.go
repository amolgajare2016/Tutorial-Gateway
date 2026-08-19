package main

import "fmt"

func main() {
	arr := [][]int{{1, 2, 3, 4}, {4, 5, 6, 7}, {8, 9, 10, 7}, {1, 4, 7, 9}}
	row := 4
	for i := 0; i < row; i++ {
		for j := 0; j < i; j++ {
			arr[i][j] = 0
		}
	}
	for i := 0; i < row; i++ {
		fmt.Println(arr[i])
	}
}

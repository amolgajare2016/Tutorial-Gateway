package main

import "fmt"

func main() {
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	k := 2
	row := 3
	for i := 0; i < row; i++ {
		for j := 0; j < row; j++ {
			arr[i][j] = arr[i][j] * k
		}
	}
	for i := 0; i < row; i++ {
		fmt.Println(arr[i])
	}
}

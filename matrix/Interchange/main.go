package main

import "fmt"

func main() {
	arr := [][]int{{1, 2, 3}, {7, 8, 9}, {2, 6, 9}}
	row := 3
	for i := 0; i < 3; i++ {
		fmt.Println(arr[i])
	}

	for i := 0; i < row; i++ {
		temp := arr[i][i]
		arr[i][i] = arr[i][row-i-1]
		arr[i][row-i-1] = temp
	}
	fmt.Println("==================================================================\n")
	for i := 0; i < 3; i++ {
		fmt.Println(arr[i])
	}
}

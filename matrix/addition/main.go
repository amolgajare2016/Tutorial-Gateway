package main

import "fmt"

func main() {
	arr1 := make([][]int, 0, 2)
	arr2 := make([][]int, 0, 2)
	additon := make([][]int, 2)
	arr1 = [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	arr2 = [][]int{{12, 14}, {16, 18}, {20, 22}, {24, 26}}
	fmt.Println("arr1 :", arr1)

	fmt.Println("arr2 :", arr2)

	for i := 0; i < 2; i++ {
		additon[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			additon[i][j] = arr1[i][j] + arr2[i][j]
		}
	}
	fmt.Println("addition :", additon)
}

package main

import "fmt"

func main() {
	arr1 := make([][]int, 2)
	arr2 := make([][]int, 2)
	add := make([][]int, 2)
	sub := make([][]int, 2)
	mul := make([][]int, 2)
	div := make([][]int, 2)
	mod := make([][]int, 2)
	arr1 = [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	arr2 = [][]int{{9, 8}, {7, 6}, {5, 4}, {3, 2}}
	for i := 0; i < 2; i++ {
		add[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			add[i][j] = arr1[i][j] + arr2[i][j]
		}
	}
	for i := 0; i < 2; i++ {
		sub[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			sub[i][j] = arr1[i][j] - arr2[i][j]
		}
	}
	for i := 0; i < 2; i++ {
		mul[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			mul[i][j] = arr1[i][j] * arr2[i][j]
		}
	}
	for i := 0; i < 2; i++ {
		div[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			div[i][j] = arr2[i][j] / arr1[i][j]
		}
	}
	for i := 0; i < 2; i++ {
		mod[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			mod[i][j] = arr1[i][j] % arr2[i][j]
		}
	}
	fmt.Println("addition :\t", add)
	fmt.Println("subtraction :\t", sub)
	fmt.Println("multiplication :\t", mul)
	fmt.Println("division :\t", div)
	fmt.Println("modolo :\t", mod)
}

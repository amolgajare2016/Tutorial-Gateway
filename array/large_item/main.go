package main

import "fmt"

func main() {
	arr := []int{1, 5, 9, 100, 11, 65, 32, 45}
	max := 0
	for _, v := range arr {
		if max < v {
			max = v
		}
	}
	fmt.Println("largest element is :", max)
}

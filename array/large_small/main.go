package main

import "fmt"

func main() {
	arr := []int{13, 4, 25, 63, 11, 78, 1, 99, 12, 48, 100, 77}
	small := arr[0]
	large := 0
	for _, v := range arr {
		if small > v {
			small = v
		}
		if large < v {
			large = v
		}
	}
	fmt.Println("smallet element is :", small)
	fmt.Println("largest element is  :", large)
}

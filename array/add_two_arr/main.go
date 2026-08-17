package main

import "fmt"

func main() {
	oddArr := []int{1, 3, 5, 7, 9}
	evenArr := []int{2, 4, 6, 8, 10}
	length := (len(oddArr) + len(evenArr)) / 2
	addArr := make([]int, length)
	for i := 0; i < length; i++ {
		addArr[i] = oddArr[i] + evenArr[i]
	}
	for _, v := range addArr {
		fmt.Println(v)
	}
}

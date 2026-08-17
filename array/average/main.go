package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 5, 8, 9, 6}
	length := len(s1)
	res := 0
	for _, v := range s1 {
		res += v
	}
	fmt.Println("res :", res)
	fmt.Println("length :", length)
	fmt.Println("average is :", res/length)
}

package main

import "fmt"

func main() {
	s1 := []int{1, 5, 9, 2, 3, 3, 3, 4, 4, 1, 4, 6, 7, 3, 2, 1}
	m := make(map[int]int)
	for _, v := range s1 {
		m[v]++
	}
	for k, v := range m {
		if v > 1 {
			fmt.Println(k, " is repeated ", v, "times ")
		} else {
			fmt.Println(k, " not repeaded")
		}
	}
}

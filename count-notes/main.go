package main

import "fmt"

func main() {
	notes := []int{500, 100, 50, 20, 10, 5, 2, 1}
	amount := 145
	temp := amount
	m1 := make(map[int]int)
	for i := 0; i < len(notes); i++ {
		if notes[i] <= temp {
			count := temp / notes[i]
			temp = temp % notes[i]
			m1[notes[i]] = count
		}
	}
	fmt.Println("m1 :", m1)
}

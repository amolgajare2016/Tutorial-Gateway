package main

import "fmt"

func main() {
	s1 := []int{2, 4, 6, 8, 10}
	s2 := []int{1, 3, 5, 7, 9}
	length := (len(s1) + len(s2)) / 2
	s3 := make([]int, 0, length)
	for i := 0; i < length; i++ {
		s3 = append(s3, s1[i]*s2[i])
	}
	fmt.Println("s3:", s3)
	for i := 0; i < length; i++ {
		fmt.Print("\t", s3[i])
		fmt.Println("")
	}
}

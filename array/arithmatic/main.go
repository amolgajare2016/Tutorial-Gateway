package main

import "fmt"

func main() {
	s1 := []int{1, 3, 5, 7, 9}
	s2 := []int{2, 4, 6, 8, 10}
	length := (len(s1) + len(s2)) / 2
	for i := 0; i < length; i++ {
		fmt.Print("\taddition :", s1[i]+s2[i])
		fmt.Print("\tsubtraction :", s1[i]-s2[i])
		fmt.Print("\tmultiplication :", s1[i]*s2[i])
		fmt.Print("\tdivision :", s1[i]/s2[i])
		fmt.Println("")
	}
}

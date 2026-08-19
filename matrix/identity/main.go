package main

import "fmt"

func main() {
	arr := [][]int{{1, 0}, {0, 1}}
	flag := false
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if i == j {
				if arr[i][j] != 1 {
					flag = true
				}
			}
			if i != j {
				if arr[i][j] != 0 {
					flag = true
				}
			}
		}
	}
	if flag {
		fmt.Println("mattrix is not identity matrix")
	} else {
		fmt.Println("matrix is identity matrix")
	}
}

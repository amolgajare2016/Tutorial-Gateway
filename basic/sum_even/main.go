package main

import "fmt"

func main() {
	num := 10
	res := 0
	for v := range num {
		if v%2 == 0 {
			res += v
		}
	}
	fmt.Println("RES", res)
}

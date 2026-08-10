package main

import "fmt"

func main() {
	n := 1
	count(n)
}

func count(n int) {
	if n < 101 {
		fmt.Println(n)
		count(n + 1)
	}
}

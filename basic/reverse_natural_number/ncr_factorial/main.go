package main

import "fmt"

func main() {
	var n, r int
	n = 10
	r = 5
	res := factorial(5)
	t := res
	fmt.Println("t;", t)
	c := factorial(n) / (factorial(r) * factorial(n-r))
	fmt.Println("result is :", c)
}

func factorial(n int) int {
	var res=1
	for i := 1; i <= n; i++ {
		res = res * i
	}
	return res
}

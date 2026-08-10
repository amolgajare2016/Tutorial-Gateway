package main

import "fmt"

func main() {
	num := 1477415
	temp := num
	rem := 0
	res := 0
	for temp > 0 {
		rem = temp % 10
		res = res*10 + rem
		temp = temp / 10
	}
	if res == num {
		fmt.Println("number is palindrom")
	} else {
		fmt.Println("number is not palindrom")
	}
}

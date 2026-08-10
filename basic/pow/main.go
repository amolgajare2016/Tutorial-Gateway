package main

import (
	"fmt"
	"math"
)

func main() {
	var nub, expo float64
	nub = 10
	expo = 2
	res := math.Pow(nub, expo)
	fmt.Println("result :", res)
}

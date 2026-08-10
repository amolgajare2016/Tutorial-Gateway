package main

import "fmt"

func main() {
	var cost, sold float64
	cost = 1000
	sold = 120
	if cost > sold {
		fmt.Println("loss :", cost-sold)
	} else if cost < sold {
		fmt.Println("profit :", sold-cost)
	} else {
		fmt.Println("no profit no loss")
	}
}

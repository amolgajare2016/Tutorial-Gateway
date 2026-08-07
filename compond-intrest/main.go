package main

import (
	"fmt"
	"math"
)

func main() {
	var principal float64
	var intrest float64
	var compoundIntrest float64
	var tenure float64
	fmt.Println("Enter principa; amount")
	fmt.Scanln(&principal)
	fmt.Println("Enter rate of intrest")
	fmt.Scanln(&intrest)
	fmt.Println("enter tenure")
	fmt.Scanln(&tenure)
	total := principal * math.Pow(1+intrest/100, tenure)
	compoundIntrest = total - principal
	fmt.Println("total :", total)
	fmt.Println("compound intrest :", compoundIntrest)
}

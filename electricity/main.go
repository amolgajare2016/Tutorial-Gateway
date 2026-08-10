package main

import "fmt"

func main() {
	var billAmount, unit, superCharge float64
	unit = 175
	if unit >= 50 {
		billAmount = unit * 2.62
		superCharge = 25
	} else if unit <= 100 {
		billAmount = 130 + (unit-50)*3.25
		superCharge = 35
	} else if unit <= 200 {
		billAmount = 130 + 162.50 + (unit-100)*5.65
		superCharge = 45
	} else {
		billAmount = 130 + 162.50 + 526 + (unit-200)*7.75
		superCharge = 55
	}
	fmt.Println("billAmount :", billAmount+superCharge)
}

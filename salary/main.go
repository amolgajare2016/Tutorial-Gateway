package main

import "fmt"

func main() {
	var basicSalary, hra, da, grossalary float64
	fmt.Println("enter salary")
	fmt.Scanln(&basicSalary)
	if basicSalary <= 10000 {
		hra = (basicSalary * 8) / 100
		da = (basicSalary * 10) / 100
	} else if basicSalary <= 20000 {
		hra = (basicSalary * 16) / 100
		da = (basicSalary * 20) / 100
	} else {
		hra = (basicSalary * 24) / 100
		da = (basicSalary * 30) / 100
	}
	grossalary = basicSalary + da + hra
	fmt.Println("your gross salary is :", grossalary)
}

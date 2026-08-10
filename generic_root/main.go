package main

func main() {
	num := 145254
	for num > 10 {
		temp := general(num)
		num = temp
	}
}

func general(temp int) int {
	var rem, res int
	for temp > 0 {
		rem = temp % 10
		res = res + rem
		temp = temp / 10
	}
	return res
}

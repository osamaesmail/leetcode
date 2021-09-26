package main

import (
	"math"
)

func main()  {
	//fmt.Println(reverse(123))
	//fmt.Println(reverse(-123))
	//fmt.Println(reverse(120))
	//fmt.Println(reverse(0))
	//fmt.Println(reverse(1534236469))
	//fmt.Println(reverse(-2147483412))
	//fmt.Println(reverse(1463847412))
}

// Pop and Push Digits solution
func reverse(x int) int {
	var rev int
	for x != 0 {
		pop := x % 10
		rev = rev * 10 + pop
		x /= 10
	}
	if rev > math.MaxInt32 || rev < math.MinInt32 {
		return 0
	}
	return rev
}
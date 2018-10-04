package main

import (
	"fmt"
)

func reverse(x int) int {
	maxInt := 2
	for i := 2; i <= 31; i++ {
		maxInt *= 2
	}

	mult := 1
	if x < 0 {
		x *= -1
		mult = -1
	}
	var r int
	for x >= 10 {
		r = r*10 + x%10
		x /= 10
	}
	r = r*10 + x
	if mult > 0 && r > maxInt-1 || mult < 0 && r > 1*maxInt {
		return 0
	}
	return r * mult
}

func main() {
	a := 123456789
	fmt.Println(reverse(a))
}

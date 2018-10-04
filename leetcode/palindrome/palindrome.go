package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	}

	var s []int
	var prev, digit int
	for x >= 10 {
		prev = x
		x = x / 10
		digit = prev - x*10
		s = append(s, digit)
	}
	s = append(s, x)

	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}
	return true
}

func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	}

	var b int
	for b < x {
		b = b*10 + x%10
		x = x / 10
	}

	if b == x || b/10 == x {
		return true
	}
	return false
}

func main() {
	a := 1001
	if isPalindrome2(a) {
		fmt.Printf("%d is palindrome\n", a)
	} else {
		fmt.Printf("%d is not palindrome\n", a)
	}
}

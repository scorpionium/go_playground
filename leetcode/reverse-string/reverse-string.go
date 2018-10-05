package main

import (
	"fmt"
)

func reverseString(s string) string {
	n := len(s)
	for i := 0; i < n/2; i++ {
		tmp1 := s[i]
		tmp2 := s[n-i-1]
		s = s[:i] + string(tmp2) + s[i+1:n-i-1] + string(tmp1) + s[n-i:]

	}
	return s
}

func main() {
	input := "Hello World"
	fmt.Println(reverseString(input))
}

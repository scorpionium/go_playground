package main

import (
	"fmt"
)

func toLowerCase(str string) string {
	for i, char := range str {
		if char >= 'A' && char <= 'Z' {
			replacement := char + 32
			str = str[:i] + string(replacement) + str[i+1:]
		}
	}
	return str
}

func main() {
	input := "Lovely"
	fmt.Println(toLowerCase(input))
}

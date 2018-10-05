package main

import (
	"fmt"
)

type stack []rune

func (s stack) Empty() bool {
	return len(s) == 0
}

func (s stack) Peek() rune {
	return s[len(s)-1]
}

func (s *stack) Push(r rune) {
	(*s) = append((*s), r)
}

func (s *stack) Pop() rune {
	r := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return r
}

func (s stack) Print() {
	for _, v := range s {
		fmt.Print(" ", string(v))
	}
	fmt.Println("")
}

func isValid(str string) bool {
	if "" == str {
		return true
	}
	var s stack

	for _, char := range str {
		if s.Empty() {
			s.Push(char)
		} else {
			peek := s.Peek()
			switch char {
			case ')':
				if peek == '(' {
					s.Pop()
				} else {
					s.Push(char)
				}
			case ']':
				if peek == '[' {
					s.Pop()
				} else {
					s.Push(char)
				}
			case '}':
				if peek == '{' {
					s.Pop()
				} else {
					s.Push(char)
				}
			default:
				s.Push(char)
			}
		}
	}

	return s.Empty()
}

func main() {
	s := "{[]}"
	if isValid(s) {
		fmt.Println("Valid!")
	} else {
		fmt.Println("Invalid!")
	}
}

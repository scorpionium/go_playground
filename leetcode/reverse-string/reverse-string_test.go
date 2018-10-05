package main

import (
	"testing"
)

func Test1(t *testing.T) {
	input := "hello"
	e := "olleh"
	v := reverseString(input)
	if e != v {
		t.Errorf("Expected %v, got %v", e, v)
	}
}

func Test2(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	e := "amanaP :lanac a ,nalp a ,nam A"
	v := reverseString(input)
	if e != v {
		t.Errorf("Expected %v, got %v", e, v)
	}
}

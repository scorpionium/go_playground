package main

import (
	"testing"
)

func Test1(t *testing.T) {
	input := "()"
	e := true
	v := isValid(input)
	if e != v {
		t.Errorf("Expected %v, got %v", e, v)
	}
}

func Test2(t *testing.T) {
	input := "()[]{}"
	e := true
	v := isValid(input)
	if e != v {
		t.Errorf("Expected %v, got %v", e, v)
	}
}

func Test3(t *testing.T) {
	input := "(]"
	e := false
	v := isValid(input)
	if e != v {
		t.Errorf("Expected %v, got %v", e, v)
	}
}

func Test4(t *testing.T) {
	input := "([)]"
	e := false
	v := isValid(input)
	if e != v {
		t.Errorf("Expected %v, got %v", e, v)
	}
}

func Test5(t *testing.T) {
	input := "{[]}"
	e := true
	v := isValid(input)
	if e != v {
		t.Errorf("Expected %v, got %v", e, v)
	}
}

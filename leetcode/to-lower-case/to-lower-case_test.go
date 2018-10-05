package main

import (
	"testing"
)

func Test1(t *testing.T) {
	input := "Hello"
	e := "hello"
	v := toLowerCase(input)
	if e != v {
		t.Errorf("Expected %v, got %v", e, v)
	}
}

func Test2(t *testing.T) {
	input := "here"
	e := "here"
	v := toLowerCase(input)
	if e != v {
		t.Errorf("Expected %v, got %v", e, v)
	}
}

func Test3(t *testing.T) {
	input := "LOVELY"
	e := "lovely"
	v := toLowerCase(input)
	if e != v {
		t.Errorf("Expected %v, got %v", e, v)
	}
}

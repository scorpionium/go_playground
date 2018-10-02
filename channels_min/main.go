package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	LEN = 30
)

func min2int(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min(s []int, c chan int) {
	min := s[0]
	for _, v := range s[1:] {
		if v < min {
			min = v
		}
	}
	c <- min
}

func main() {
	var s []int
	s = make([]int, LEN, LEN)
	rand.Seed(13)
	for i, _ := range s {
		s[i] = rand.Int()
	}

	c := make(chan int)

	start := time.Now()
	go min(s, c)
	m := <-c
	elapsed := time.Since(start)
	fmt.Printf("Min = %d, time took = %s\n", m, elapsed)

	start = time.Now()
	go min(s[:len(s)/2], c)
	go min(s[len(s)/2:], c)
	min1, min2 := <-c, <-c
	elapsed = time.Since(start)
	fmt.Printf("Min = %d, time took = %s\n", min2int(min1, min2), elapsed)
}

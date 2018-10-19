package main

import (
	"fmt"
)

type queue []int

func (q queue) Size() int {
	return len(q)
}

func (q queue) Empty() bool {
	return q.Size() == 0
}

func (q queue) Front() int {
	return q[0]
}

func (q *queue) Push(x int) {
	(*q) = append((*q), x)
}

func (q *queue) Pop() int {
	x := (*q)[0]
	(*q) = (*q)[1:]
	return x
}

func main() {
	var q queue
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range input {
		q.Push(v)
	}
	fmt.Println(q.Front())
	fmt.Println(q.Size())
	for q.Empty() == false {
		fmt.Printf("%d ", q.Pop())
	}
	fmt.Println()
}

/*
How many unique, 7-digit phone numbers can a "chess knight" dial, starting from the "1"
1 2 3
4 5 6
7 8 9
  0

e.g. "1604061" and "1616161" are valid phone numbers.

For example, there are 5 unique 3-digit phone numbers that this knight can dial.
*/

package main

import "fmt"

type Move struct {
	c int
	r int
}

func checkField(c, r int, d [4][3]int) (bool, Move) {
	if d[r][c] != -1 {
		return true, Move{c: c, r: r}
	} else {
		return false, Move{}
	}
}

func getNextMoves(c, r int, d [4][3]int) []Move {
	maxR := 4
	maxC := 3
	res := []Move{}

	if c-1 >= 0 && r-2 >= 0 {
		if f, m := checkField(c-1, r-2, d); f == true {
			res = append(res, m)
		}
	}

	if c-2 >= 0 && r-1 >= 0 {
		if f, m := checkField(c-2, r-1, d); f == true {
			res = append(res, m)
		}
	}

	if c+1 < maxC && r-2 >= 0 {
		if f, m := checkField(c+1, r-2, d); f == true {
			res = append(res, m)
		}
	}

	if c+2 < maxC && r-1 >= 0 {
		if f, m := checkField(c+2, r-1, d); f == true {
			res = append(res, m)
		}
	}

	if c-1 >= 0 && r+2 < maxR {
		if f, m := checkField(c-1, r+2, d); f == true {
			res = append(res, m)
		}
	}

	if c-2 >= 0 && r+1 < maxR {
		if f, m := checkField(c-2, r+1, d); f == true {
			res = append(res, m)
		}
	}

	if c+1 < maxC && r+2 < maxR {
		if f, m := checkField(c+1, r+2, d); f == true {
			res = append(res, m)
		}
	}

	if c+2 < maxC && r+1 < maxR {
		if f, m := checkField(c+2, r+1, d); f == true {
			res = append(res, m)
		}
	}

	return res

}

func dial(n, c, r int, d [4][3]int) int {
	if n == 1 {
		return 1
	}

	moves := getNextMoves(c, r, d)
	var s int
	for _, v := range moves {
		s += dial(n-1, v.c, v.r, d)
	}
	return s
}

func main() {

	d := [4][3]int{
		[3]int{1, 2, 3},
		[3]int{4, 5, 6},
		[3]int{7, 8, 9},
		[3]int{-1, 0, -1},
	}

	n := 7
	c := 0
	r := 0
	fmt.Printf("Result = %d\n", dial(n, c, r, d))
}

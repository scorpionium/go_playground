package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func hourGlassSum(in [][]int32, r, c int) int32 {
	hourGlassMatrix := [][]int32{
		{1, 1, 1},
		{0, 1, 0},
		{1, 1, 1},
	}
	n := 3
	sum := int32(0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum += hourGlassMatrix[i][j] * in[i+r][j+c]
		}
	}
	return sum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(6) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}
	maxSum := int32(math.MinInt32)

	n := 6
	for i := 0; i < n-2; i++ {
		for j := 0; j < n-2; j++ {
			sum := hourGlassSum(arr, i, j)
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	fmt.Println(maxSum)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)
	c := 0
	maxc := 0
	for n > 0 {
		fmt.Printf("%d", n&1)
		if (n & 1) > 0 {
			c++
		} else {
			c = 0
		}
		n = n >> 1
		if c > maxc {
			maxc = c
		}
	}
	fmt.Println(maxc)
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

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

const file = "input.txt"

func getScanner(file string) *bufio.Scanner {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	return scanner
}

func main() {
	puzzle1()
	puzzle2()
}

func value(bs []byte) uint8 {
	i, _ := strconv.ParseInt(string(bs), 10, 8)
	return uint8(i)
}

func puzzle1() {
	scanner := getScanner(file)

	var overlapping int32 = 0
	for scanner.Scan() {
		bs := scanner.Bytes()

		firstDash, secondDash := bytes.IndexByte(bs, '-'), bytes.LastIndexByte(bs, '-')
		comma := bytes.IndexByte(bs, ',')

		oneStart, oneEnd := value(bs[0:firstDash]), value(bs[firstDash+1:comma])
		twoStart, twoEnd := value(bs[comma+1:secondDash]), value(bs[secondDash+1:])

		if (oneStart >= twoStart && oneEnd <= twoEnd) ||
			(twoStart >= oneStart && twoEnd <= oneEnd) {
			overlapping++
		}
	}
	fmt.Printf("Puzzle1 Overlapping: %v\n", overlapping)
}

func puzzle2() {
	scanner := getScanner(file)

	var overlapping int32 = 0
	for scanner.Scan() {
		bs := scanner.Bytes()

		firstDash, secondDash := bytes.IndexByte(bs, '-'), bytes.LastIndexByte(bs, '-')
		comma := bytes.IndexByte(bs, ',')

		oneStart, oneEnd := value(bs[0:firstDash]), value(bs[firstDash+1:comma])
		twoStart, twoEnd := value(bs[comma+1:secondDash]), value(bs[secondDash+1:])

		if (oneStart >= twoStart && oneStart <= twoEnd) ||
			(twoStart >= oneStart && twoStart <= oneEnd) {
			overlapping++
		}

	}
	fmt.Printf("Puzzle2 Overlapping: %v\n", overlapping)
}

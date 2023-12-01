package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"strconv"
)

func main() {
	puzzle1()
}

func getScanner(file string) *bufio.Scanner {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	return scanner
}

func puzzle1() {
	const file = "mini.txt"
	scanner := getScanner(file)

	var rope image.Point

	var visisted = map[image.Point]bool{}
	for scanner.Scan() {
		bs := scanner.Bytes()

		pos, num := bs[0], bs[2:]
		n, _ := strconv.Atoi(string(num))
		fmt.Printf("%v - %v\n", string(pos), n)

		for i := 0; i < n; i++ {
			distance := 
		}

	}
}

var points = map[byte]image.Point{'U': {0, -1}, 'D': {0, 1}, 'L': {-1, 0}, 'R': {1, 0}}

func puzzle1() {
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

const file = "input.txt"

func main() {
	puzzle1()
	puzzle2()
}

func priority(b byte) int32 {
	// A-Z = 65-90
	// a-z = 97-122
	if b >= 97 {
		// Must be a-z
		return int32(b) - 96

	}
	// Must be A-Z
	return int32(b) - 38
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
	scanner := getScanner(file)

	var priorities int32 = 0
	for scanner.Scan() {
		text := scanner.Text()

		pocketSize := len(text) / 2
		c1, c2 := text[0:pocketSize], text[pocketSize:]

		var items = map[byte]bool{}
		for i := range c1 {
			items[c1[i]] = true
		}

		for i := range c2 {
			if valid, ok := items[c2[i]]; ok && valid {
				priorities += priority(c2[i])
				items[c2[i]] = false
			}
		}
	}
	fmt.Printf("Puzzle1 Priorities: %v\n", priorities)
}

func puzzle2() {
	scanner := getScanner(file)

	var groupings uint8 = 0
	var items = map[byte]uint8{}
	var priorities int32 = 0

	for scanner.Scan() {
		text := scanner.Text()

		var lineItems = [52]bool{}
		for i := range text {
			p := priority(text[i]) - 1
			if !lineItems[p] {
				items[text[i]]++
				lineItems[p] = true
			}
		}

		groupings++
		if groupings == 3 {
			for k, v := range items {
				if v == 3 {
					priorities += priority(k)
					break
				}
			}

			groupings = 0
			items = map[byte]uint8{}
		}
	}
	fmt.Printf("Puzzle2 Priorities: %v\n", priorities)
}

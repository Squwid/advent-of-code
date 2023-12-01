package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

const file = "input.txt"

func main() {
	puzzle1()
	puzzle2()
}

type stack struct {
	nodes []*node
}

type node struct {
	value byte
}

func newStack() *stack {
	return &stack{nodes: []*node{}}
}

func (s *stack) push(n *node) {
	s.nodes = append(s.nodes, n)
}

func (s *stack) pushInitial(n *node) {
	s.nodes = append([]*node{n}, s.nodes...)
}

func (s *stack) pop() *node {
	if len(s.nodes) == 0 {
		return nil
	}
	n := s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[:len(s.nodes)-1]
	return n
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
	var stacks = map[int]*stack{}

	scanner := getScanner(file)

	var scanningCrates = true
	for scanner.Scan() {
		var openBracket = false // True when '[' was the last char parsed.

		bs := scanner.Bytes()

		if len(bs) > 0 && bs[1] == '1' {
			scanningCrates = false
			continue
		}

		if scanningCrates {
			for i := range bs {
				if openBracket {
					index := ((i - 1) / 4) + 1
					if _, ok := stacks[index]; !ok {
						stacks[index] = newStack()
					}
					stacks[index].pushInitial(&node{value: bs[i]})
					// fmt.Printf("Crate %s to %v\n", string(bs[i]), index)
				}

				openBracket = bs[i] == '['

			}
		} else {
			// Only parse pages that start with 'move' otherwise ignore
			if len(bs) == 0 || bs[0] != 'm' {
				continue
			}

			// Base indexes where 't' is since its easy
			tIndex := bytes.IndexByte(bs, 't')
			amount, _ := strconv.Atoi(string(bs[5 : bytes.IndexByte(bs, 'f')-1]))
			from, _ := strconv.Atoi(string(bs[tIndex-2 : tIndex-1]))
			to, _ := strconv.Atoi(string(bs[tIndex+3:]))
			// fmt.Printf("move %v from %v to %v\n", amount, from, to)

			for i := 0; i < amount; i++ {
				n := stacks[from].pop()
				// fmt.Printf("moving %v from %v to %v\n", string(n.value), from, to)
				stacks[to].push(n)
			}

		}
	}
	// Print top values in order
	fmt.Printf("Puzzle1: ")
	for i := 1; i < len(stacks)+1; i++ {
		fmt.Printf("%s", string(stacks[i].pop().value))
	}
	fmt.Printf("\n")
}

func puzzle2() {
	var stacks = map[int]*stack{}

	scanner := getScanner(file)

	var scanningCrates = true
	for scanner.Scan() {
		var openBracket = false // True when '[' was the last char parsed.

		bs := scanner.Bytes()

		if len(bs) > 0 && bs[1] == '1' {
			scanningCrates = false
			continue
		}

		if scanningCrates {
			for i := range bs {
				if openBracket {
					index := ((i - 1) / 4) + 1
					if _, ok := stacks[index]; !ok {
						stacks[index] = newStack()
					}
					stacks[index].pushInitial(&node{value: bs[i]})
					// fmt.Printf("Crate %s to %v\n", string(bs[i]), index)
				}

				openBracket = bs[i] == '['

			}
		} else {
			// Only parse pages that start with 'move' otherwise ignore
			if len(bs) == 0 || bs[0] != 'm' {
				continue
			}

			tIndex := bytes.IndexByte(bs, 't')
			amount, _ := strconv.Atoi(string(bs[5 : bytes.IndexByte(bs, 'f')-1]))
			from, _ := strconv.Atoi(string(bs[tIndex-2 : tIndex-1]))
			to, _ := strconv.Atoi(string(bs[tIndex+3:]))
			// fmt.Printf("move %v from %v to %v\n", amount, from, to)

			var tempStack = newStack()
			for i := 0; i < amount; i++ {
				tempStack.push(stacks[from].pop())
			}
			for i := 0; i < amount; i++ {
				stacks[to].push(tempStack.pop())
			}
		}
	}
	// Print top values in order
	fmt.Printf("Puzzle2: ")
	for i := 1; i < len(stacks)+1; i++ {
		fmt.Printf("%s", string(stacks[i].pop().value))
	}
	fmt.Printf("\n")
}

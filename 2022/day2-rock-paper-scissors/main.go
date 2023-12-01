package main

import (
	"bufio"
	"fmt"
	"os"
)

const file = "puzzle1.txt"

const (
	ROCK byte = iota
	PAPER
	SCISSORS

	WIN
	LOSS
	TIE
)

var points = map[byte]int32{
	ROCK:     1,
	PAPER:    2,
	SCISSORS: 3,

	WIN:  6,
	TIE:  3,
	LOSS: 0,
}

func main() {
	puzzle1()
	puzzle2()
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
	var outcomes = map[byte]byte{
		(ROCK << 3) | PAPER:     WIN,
		(ROCK << 3) | SCISSORS:  LOSS,
		(PAPER << 3) | SCISSORS: WIN,
		(PAPER << 3) | ROCK:     LOSS,
		(SCISSORS << 3) | ROCK:  WIN,
		(SCISSORS << 3) | PAPER: LOSS,
	}

	var play = map[byte]byte{
		'A': ROCK,
		'X': ROCK,
		'B': PAPER,
		'Y': PAPER,
		'C': SCISSORS,
		'Z': SCISSORS,
	}

	scanner := getScanner(file)

	var score int32 = 0
	for scanner.Scan() {
		text := scanner.Text()
		op, me := text[0], text[2]

		if play[op] == play[me] {
			score += points[TIE]
		} else {
			outcome := outcomes[play[op]<<3|play[me]]
			score += points[outcome]
		}
		score += points[play[me]]
	}
	fmt.Printf("Puzzle1 score: %v\n", score)
}

func puzzle2() {
	var outcomes = map[byte]byte{
		(ROCK << 3) | LOSS:     SCISSORS,
		(ROCK << 3) | WIN:      PAPER,
		(PAPER << 3) | LOSS:    ROCK,
		(PAPER << 3) | WIN:     SCISSORS,
		(SCISSORS << 3) | LOSS: PAPER,
		(SCISSORS << 3) | WIN:  ROCK,
	}

	var play = map[byte]byte{
		'A': ROCK,
		'B': PAPER,
		'C': SCISSORS,

		'X': LOSS,
		'Y': TIE,
		'Z': WIN,
	}

	scanner := getScanner(file)

	var score int32 = 0
	for scanner.Scan() {
		text := scanner.Text()
		op, outcome := text[0], text[2]

		// What i need to throw to get the outcome
		var throw byte

		if play[outcome] == TIE {
			score += points[TIE]
			throw = play[op]
		} else {
			score += points[play[outcome]]
			throw = outcomes[play[op]<<3|play[outcome]]
		}
		score += points[throw]
	}
	fmt.Printf("Puzzle2 score: %v\n", score)
}

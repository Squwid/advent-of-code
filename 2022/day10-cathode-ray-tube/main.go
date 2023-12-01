package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

var instructions = map[string]int32{
	"noop": 1,
	"addx": 2,
}

func puzzle1() {
	scanner := getScanner(file)

	var cycle, register int32 = 1, 1

	// Map cycle to register value to see outputs
	var debugger = map[int32]int32{}

	for scanner.Scan() {
		text := scanner.Text()
		args := strings.Split(text, " ")

		instruction := args[0]
		for i := int32(0); i < instructions[instruction]; i++ {
			debugger[cycle] = register
			cycle++
		}

		if instruction == "addx" {
			i, _ := strconv.Atoi(args[1])
			register += int32(i)
		}
	}

	var solution int32 = 0
	for i := int32(0); i <= 240; i++ {
		if (i+20)%40 == 0 {
			solution += i * debugger[i]
		}
	}

	fmt.Printf("Puzzle 1 answer: %v\n", solution)
}

func puzzle2() {
	scanner := getScanner(file)

	var cycle, register int32 = 1, 1
	var output = [6][40]rune{}

	// Map cycle to register value to see outputs
	populate := func(index, register int32) bool {
		diff := register - index
		return diff >= -1 && diff <= 1
	}

	for scanner.Scan() {
		text := scanner.Text()
		args := strings.Split(text, " ")

		instruction := args[0]
		for i := int32(0); i < instructions[instruction]; i++ {
			p := populate((cycle-1)%40, register)
			if p {
				output[(cycle-1)/40][cycle%40] = '🎁'
			} else {
				output[(cycle-1)/40][cycle%40] = '🎄'
			}
			cycle++
		}

		if instruction == "addx" {
			i, _ := strconv.Atoi(args[1])
			register += int32(i)
		}
	}

	fmt.Println("Puzzle 2 Output:")
	for i := 0; i < len(output); i++ {
		fmt.Println(string(output[i][:]))
	}
}

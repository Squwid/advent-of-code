package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const file = "puzzle1.txt"
const elfLeaderboardSize = 3

func main() {
	puzzle1()
	puzzle2()
}

type elf struct {
	index    int32
	calories int32
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

	var elfMax *elf
	var elfIndex int32 = 1
	var calories int32 = 0

	compareMax := func(index, calories int32) {
		var e = elf{
			index:    elfIndex,
			calories: calories,
		}
		if elfMax == nil || elfMax.calories < e.calories {
			elfMax = &e
		}
	}

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			compareMax(elfIndex, calories)

			calories = 0
			elfIndex++
		} else {
			i, _ := strconv.ParseInt(text, 10, 32)
			calories += int32(i)
		}
	}
	if calories != 0 {
		compareMax(elfIndex, calories)
	}

	// Elf 151 has the most calories of 69912
	fmt.Printf("Elf %v has the most calories of %v\n", elfMax.index, elfMax.calories)
}

type elfLeaderboard [elfLeaderboardSize]*elf

func (lb *elfLeaderboard) compare(e *elf, index int) {
	for i := index; i < len(lb); i++ {
		if lb[i] == nil {
			lb[i] = e
			return
		}

		if lb[i].calories < e.calories {
			removalElf := lb[i]
			lb[i] = e
			lb.compare(removalElf, i+1)
			return
		}
	}
}

// helper to print leaderboard
func (lb *elfLeaderboard) printlb() {
	for i := 0; i < len(lb); i++ {
		fmt.Printf("%v: ", i)
		if lb[i] == nil {
			fmt.Printf("nil")
		} else {
			fmt.Printf("(%v) - %v", lb[i].index, lb[i].calories)
		}
		fmt.Printf(", ")
	}
	fmt.Printf("\n")
}

func (lb *elfLeaderboard) calories() (cals int32) {
	for i := 0; i < len(lb); i++ {
		if lb[i] == nil {
			continue
		}
		cals += lb[i].calories
	}
	return
}

func puzzle2() {
	scanner := getScanner(file)

	var elfLb = new(elfLeaderboard)
	var elfIndex int32 = 1
	var calories int32 = 0

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			elfLb.compare(&elf{
				index:    elfIndex,
				calories: calories,
			}, 0)

			calories = 0
			elfIndex++
		} else {
			i, _ := strconv.ParseInt(text, 10, 32)
			calories += int32(i)
		}
	}
	if calories != 0 {
		elfLb.compare(&elf{index: elfIndex, calories: calories}, 0)
	}

	// The top 3 elves have 208180 calories
	fmt.Printf("The top 3 elves have %v calories\n", elfLb.calories())
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

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
	scanner := getScanner("input.txt")

	var total int
	for scanner.Scan() {
		var nums []byte
		line := scanner.Bytes()

		for i := 0; i < len(line); i++ {
			if line[i] >= 49 && line[i] <= 57 {
				nums = append(nums, line[i])
			} else {
				b := convertBytes(line[i:])
				if b != 0 {
					nums = append(nums, b)
				}
			}
		}

		var num []byte
		if len(nums) == 0 {
			continue
		} else {
			num = append(num, nums[0])           // add first nunber
			num = append(num, nums[len(nums)-1]) // add last number
		}

		i, _ := strconv.Atoi(string(num))
		total += i
	}
	fmt.Println(total)
}

// convertBytes get passed the remaining bytes and checks if the next x bytes
// are a spelled out number.
func convertBytes(bs []byte) byte {
	if len(bs) >= 3 {
		if bytes.Equal(bs[0:3], []byte("one")) {
			return byte('1')
		}
		if bytes.Equal(bs[0:3], []byte("two")) {
			return byte('2')
		}
		if bytes.Equal(bs[0:3], []byte("six")) {
			return byte('6')
		}
	}
	if len(bs) >= 4 {
		if bytes.Equal(bs[0:4], []byte("four")) {
			return byte('4')
		}
		if bytes.Equal(bs[0:4], []byte("five")) {
			return byte('5')
		}
		if bytes.Equal(bs[0:4], []byte("nine")) {
			return byte('9')
		}
	}
	if len(bs) >= 5 {
		if bytes.Equal(bs[0:5], []byte("three")) {
			return byte('3')
		}
		if bytes.Equal(bs[0:5], []byte("seven")) {
			return byte('7')
		}
		if bytes.Equal(bs[0:5], []byte("eight")) {
			return byte('8')
		}
	}
	return 0
}

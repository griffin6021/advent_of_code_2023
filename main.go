package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer f.Close()

	var sum int
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		first := findFirstNumber(line)
		last := findLastNumber(line)
		fmt.Printf("line: %s, first: %s, last: %s, num: %s\n", line, first, last, first+last)
		val, err := strconv.Atoi(first + last)
		if err != nil {
			log.Fatalf("atoi failed: %v", err)
		}
		sum += val
	}
	fmt.Printf("Total: %d\n", sum)
}

// findFirstNumber finds the first number in a string by
// looking from left to right for a number.
func findFirstNumber(line string) string {
	for i := 0; i < len(line); i++ {
		if isNumber(line[i]) {
			return string([]byte{line[i]})
		}
	}
	return ""
}

// findLastNumber finds the last number in a string by
// looking from right to left for a number.
func findLastNumber(line string) string {
	for i := len(line) - 1; i >= 0; i-- {
		if isNumber(line[i]) {
			return string([]byte{line[i]})
		}
	}
	return ""
}

func isNumber(c byte) bool {
	return c >= '0' && c <= '9'
}

//look for all of the first digits
//look for all of the second digits
//combine them to make a single 2 digit number
//add all of these numbers

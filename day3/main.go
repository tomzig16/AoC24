package main

import (
	"adventofcode/shared"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines, err := shared.ReadLines("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input %v", err)
	}

	singleLine := strings.Join(lines[:], "\n")

	Part1(singleLine)
	Part2(singleLine)

}

func Part1(input string) {
	r, _ := regexp.Compile(`mul\((\d+),(\d+)\)`)
	matches := r.FindAllStringSubmatch(input, -1)

	finalSum := 0

	for _, match := range matches {
		firstNumber, _ := strconv.Atoi(match[1])
		secondNumber, _ := strconv.Atoi(match[2])
		finalSum += (firstNumber * secondNumber)
	}

	fmt.Println("Result Part 1: ", finalSum)
}

func Part2(input string) {
	r, _ := regexp.Compile(`do\(\)|don't\(\)|mul\((\d+),(\d+)\)`)
	matches := r.FindAllStringSubmatch(input, -1)

	finalSum := 0
	shouldAdd := true
	for _, match := range matches {
		switch match[0] {
		case "do()":
			shouldAdd = true
		case "don't()":
			shouldAdd = false
		default:
			if shouldAdd {
				firstNumber, _ := strconv.Atoi(match[1])
				secondNumber, _ := strconv.Atoi(match[2])
				finalSum += (firstNumber * secondNumber)
			}
		}

	}

	fmt.Println("Result Part 1: ", finalSum)
}

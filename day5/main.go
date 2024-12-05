package main

import (
	"adventofcode/shared"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	lines, err := shared.ReadLines("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input %v", err)
	}

	orderingRules, pageNumbers := GetOrderingRulesAndPageNumbers(lines)

	Part1(orderingRules, pageNumbers)
	Part2(orderingRules, pageNumbers)
	// Part2(singleLine)

}

// func GetOrderingRulesAndPageNumbers(input []string) (map[int][]int, [][]int) {
// 	orderingRules := make(map[int][]int)
// 	var pageNumbers [][]int

// 	readingOrderingRules := true
// 	pageNumbersIndex := 0
// 	for _, line := range input {
// 		if line == "" {
// 			readingOrderingRules = false
// 			continue
// 		}

// 		if readingOrderingRules {
// 			numbers, _ := shared.ConvertStringToInts(strings.Split(line, "|"))
// 			// arr, ok := orderingRules[numbers[0]]
// 			// if !ok {
// 			orderingRules[numbers[0]] = append(orderingRules[numbers[0]], numbers[1])
// 			// }
// 		} else {
// 			numbers, _ := shared.ConvertStringToInts(strings.Split(line, ","))
// 			pageNumbers = append(pageNumbers, make([]int, 0))
// 			pageNumbers[pageNumbersIndex] = append(pageNumbers[pageNumbersIndex], numbers...)
// 			pageNumbersIndex++
// 		}
// 	}
// 	return orderingRules, pageNumbers
// }

func GetOrderingRulesAndPageNumbers(input []string) (map[string]byte, [][]string) {
	orderingRules := make(map[string]byte)
	var pageNumbers [][]string

	readingOrderingRules := true
	pageNumbersIndex := 0
	for _, line := range input {
		if line == "" {
			readingOrderingRules = false
			continue
		}

		if readingOrderingRules {
			// arr, ok := orderingRules[numbers[0]]
			// if !ok {
			orderingRules[line] = 0
			// }
		} else {
			numbers := strings.Split(line, ",")
			pageNumbers = append(pageNumbers, make([]string, 0))
			pageNumbers[pageNumbersIndex] = append(pageNumbers[pageNumbersIndex], numbers...)
			pageNumbersIndex++
		}
	}
	return orderingRules, pageNumbers
}

func Part1(orderingRules map[string]byte, pageNumbers [][]string) {
	finalSum := 0
	for _, line := range pageNumbers {
		isLineCorrect := true
		for i, el := range line {
			if i == len(line)-1 {
				break
				// Last element is always correct
			}
			for j := i + 1; j < len(line); j++ {
				_, ok := orderingRules[el+"|"+line[j]]
				if !ok {
					isLineCorrect = false
					break
				}
			}
			if !isLineCorrect {
				break
			}
		}
		if isLineCorrect {
			index := (len(line) / 2)
			nToAdd, _ := strconv.Atoi(line[index])
			finalSum += nToAdd
		}
	}

	fmt.Println("Result Part 1: ", finalSum)
}

func Part2(orderingRules map[string]byte, pageNumbers [][]string) {
	finalSum := 0
	for _, line := range pageNumbers {
		isLineCorrect, failedIndex := CheckIfSequenceIsCorrect(orderingRules, line)
		if !isLineCorrect {
			index := (len(line) / 2)
			nToAdd, _ := strconv.Atoi(line[index])
			finalSum += nToAdd
			println(failedIndex)
		}
	}

	fmt.Println("Result Part 2: ")
}

// if fails then returns index of a first failed number
func CheckIfSequenceIsCorrect(orderingRules map[string]byte, line []string) (bool, int) {
	isLineCorrect := true
	failedIndex := -1
	for i, el := range line {
		if i == len(line)-1 {
			break
			// Last element is always correct
		}
		for j := i + 1; j < len(line); j++ {
			_, ok := orderingRules[el+"|"+line[j]]
			if !ok {
				isLineCorrect = false
				failedIndex = j
				break
			}
		}
		if !isLineCorrect {
			break
		}
	}
	return isLineCorrect, failedIndex
}

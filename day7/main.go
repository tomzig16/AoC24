package main

import (
	"adventofcode/shared"
	"fmt"
	"log"
	"math"
	"strings"
)

func main() {
	lines, err := shared.ReadLines("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input %v", err)
	}
	numbers := make([][]int, 0)
	for _, line := range lines {
		l := strings.Split(strings.Replace(line, ":", "", 1), " ")
		n, _ := shared.ConvertStringToInts(l)
		numbers = append(numbers, n)
	}

	Part1(numbers)
	// Part2(lines, startingPositionX, startingPositionY)

}

func Part1(input [][]int) {
	finalSum := 0
	for _, inputNumbers := range input {
		target := inputNumbers[0]
		rightSideNumbers := inputNumbers[1:]
		nOfOperators := len(rightSideNumbers) - 1

		counter := 0
		numberOfIterations := powInt(2, nOfOperators)
		for i := 0; i < numberOfIterations; i++ {
			sum := rightSideNumbers[0]
			st := "actions: "
			for x := 0; x < nOfOperators && sum < target; x++ {
				bitToCheck := 1 << x
				if (counter & bitToCheck) == 0 {
					sum += rightSideNumbers[x+1]
					st += "+"
				} else {
					sum *= rightSideNumbers[x+1]
					st += "*"
				}
			}
			if sum == target {
				finalSum += target
				break
			}
			counter++
		}
	}

	fmt.Println("Result Part 1: ", finalSum)
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func Part2(input [][]int) {
	finalSum := 0
	for _, inputNumbers := range input {
		target := inputNumbers[0]
		rightSideNumbers := inputNumbers[1:]
		operators := make([]rune, len(rightSideNumbers)-1)
		for i := 0; i < len(operators); i++ {
			operators[i] = '+'
		}

		counter := 0
		numberOfIterations := powInt(3, len(operators))
		for i := 0; i < numberOfIterations; i++ {
			sum := rightSideNumbers[0]
			st := "actions: "
			for x := 0; x < len(operators) && sum < target; x++ {
				bitToCheck := 1 << x
				if (counter & bitToCheck) == 0 {
					sum += rightSideNumbers[x+1]
					st += "+"
				} else {
					sum *= rightSideNumbers[x+1]
					st += "*"
				}
			}
			if sum == target {
				finalSum += target
				break
			}
			counter++
		}
	}

	fmt.Println("Result Part 1: ", finalSum)
}

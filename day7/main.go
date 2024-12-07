package main

import (
	"adventofcode/shared"
	"fmt"
	"log"
	"math"
	"strconv"
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

	// Part1(numbers)
	Part1Recursive(numbers)
	Part2Recursive(numbers)
	// Part2(lines, startingPositionX, startingPositionY)

}

// Old part 1 solution that broke when part 2 happened due to no good solution for base 3 numbers
// Leaving for grandchildren to cringe in the future
func Part1(input [][]int) {
	finalSum := 0
	for _, inputNumbers := range input {
		target := inputNumbers[0]
		rightSideNumbers := inputNumbers[1:]
		nOfOperators := len(rightSideNumbers) - 1

		counter := 0
		numberOfIterations := powInt(2, nOfOperators)
		fmt.Println("Target ", target)
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
			fmt.Println(strconv.Itoa(i)+": ", st)
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

// Recursive method
// func AddFunc(target int, rightSide []int, nextIndex int) int {
// 	tempResult := 0
// 	if nextIndex < len(rightSide) {
// 		tempResult = AddFunc(target, rightSide, nextIndex+1)
// 		if tempResult == target {
// 			return tempResult
// 		}
// 		tempResult = MulFunc(target, rightSide, nextIndex+1)
// 		if tempResult == target {
// 			return tempResult
// 		}
// 	}
// 	return tempResult
// }

func GetResult(target int, rightSide []int, tempTotal int, operator string, isPart2 bool) bool {
	var equationTotal int
	switch operator {
	case "+":
		equationTotal = tempTotal + rightSide[0]
		break
	case "*":
		equationTotal = tempTotal * rightSide[0]
		break
	case "||":
		tempTotalAsStr := strconv.Itoa(tempTotal)
		rightSideAsStr := strconv.Itoa(rightSide[0])
		combinedAsStr := tempTotalAsStr + rightSideAsStr
		equationTotal, _ = strconv.Atoi(combinedAsStr)
		break
	}

	if len(rightSide) == 1 {
		return equationTotal == target
	}

	sumResult := GetResult(target, rightSide[1:], equationTotal, "+", isPart2)
	mulResult := GetResult(target, rightSide[1:], equationTotal, "*", isPart2)
	concatResult := false
	if isPart2 {
		concatResult = GetResult(target, rightSide[1:], equationTotal, "||", isPart2)
	}
	return sumResult || mulResult || concatResult
}

func Part1Recursive(input [][]int) {
	finalSum := 0
	for _, inputNumbers := range input {
		target := inputNumbers[0]
		rightSideNumbers := inputNumbers[1:]

		if GetResult(target, rightSideNumbers, 0, "+", false) ||
			GetResult(target, rightSideNumbers, 0, "*", false) {
			finalSum += target
		}

	}
	fmt.Println("Result Part 1: ", finalSum)
}

func Part2Recursive(input [][]int) {
	finalSum := 0
	for _, inputNumbers := range input {
		target := inputNumbers[0]
		rightSideNumbers := inputNumbers[1:]

		if GetResult(target, rightSideNumbers, 0, "+", true) ||
			GetResult(target, rightSideNumbers, 0, "*", true) ||
			GetResult(target, rightSideNumbers, 0, "||", true) {
			finalSum += target
		}

	}
	fmt.Println("Result Part 2: ", finalSum)
}

package main

import (
	"adventofcode/shared"
	"fmt"
	"log"
	"slices"
	"strconv"
)

type Direction struct {
	dirX int
	dirY int
}

var directions map[int]Direction

func main() {
	lines, err := shared.ReadLines("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input %v", err)
	}
	startingPositionX, startingPositionY := FindStartingPosition(lines)
	Part1(lines, startingPositionX, startingPositionY)
	Part2(lines, startingPositionX, startingPositionY)

}

func FindStartingPosition(lines []string) (int, int) {
	for i, _ := range lines {
		for j, _ := range lines[i] {
			if lines[i][j] == '^' {
				return i, j
			}
		}
	}
	return -1, -1
}

func Part1(input []string, startX int, startY int) {
	var visitedCells map[string]byte = make(map[string]byte)
	directions := make(map[int]Direction)
	directions[0] = Direction{dirX: 0, dirY: -1}
	directions[1] = Direction{dirX: 1, dirY: 0}
	directions[2] = Direction{dirX: 0, dirY: 1}
	directions[3] = Direction{dirX: -1, dirY: 0}

	wallSymbol := '#'

	isOver := false
	currentDirection := 0
	verticalPos := startX
	horizontalPos := startY
	currentSymbol := input[verticalPos][horizontalPos]
	for !isOver {
		for currentSymbol != byte(wallSymbol) {
			visitedCells[PositionToKey(verticalPos, horizontalPos)] = 0
			verticalPos += directions[currentDirection].dirY
			horizontalPos += directions[currentDirection].dirX
			isPossible := IsCheckPossible(verticalPos, horizontalPos, len(input), len(input[0]))
			if !isPossible {
				isOver = true
				break
			}
			currentSymbol = input[verticalPos][horizontalPos]
		}
		if isOver {
			// We are out of the maze
			break
		}

		if currentSymbol == byte(wallSymbol) {
			verticalPos += -directions[currentDirection].dirY
			horizontalPos += -directions[currentDirection].dirX
			currentSymbol = input[verticalPos][horizontalPos]
			if currentDirection == 3 {
				currentDirection = 0
			} else {
				currentDirection++
			}
		}

	}

	fmt.Println("Result Part 1: ", len(visitedCells))
}

// So far seems incorrect
func Part2(input []string, startX int, startY int) {
	var visitedCells map[int]map[int][]int = make(map[int]map[int][]int)
	var placedObstacles map[string]byte = make(map[string]byte)
	directions = make(map[int]Direction)
	directions[0] = Direction{dirX: 0, dirY: -1}
	directions[1] = Direction{dirX: 1, dirY: 0}
	directions[2] = Direction{dirX: 0, dirY: 1}
	directions[3] = Direction{dirX: -1, dirY: 0}

	wallSymbol := '#'

	isOver := false
	currentDirection := 0
	verticalPos := startX
	horizontalPos := startY
	currentSymbol := input[verticalPos][horizontalPos]
	for !isOver {
		for currentSymbol != byte(wallSymbol) {
			el, ok := visitedCells[verticalPos][horizontalPos]
			// if we already visited the cell we can attempt placing and obstacle next to it
			if !ok {
				_, horizontalOk := visitedCells[verticalPos]
				if !horizontalOk {
					visitedCells[verticalPos] = make(map[int][]int)
				}
				visitedCells[verticalPos][horizontalPos] = make([]int, 1)
				visitedCells[verticalPos][horizontalPos][0] = currentDirection

			} else {
				if !slices.Contains(el, currentDirection) {
					visitedCells[verticalPos][horizontalPos] = append(el, currentDirection)
				}

			}

			// Check if it makes sense to place an obstacle in front
			nextDir := GetNextDir(currentDirection)

			// check horizontal
			if nextDir == 0 || nextDir == 2 {
				for _, el := range visitedCells {
					horizontalEl, ok := el[horizontalPos]
					if ok && slices.Contains(horizontalEl, nextDir) {
						obstaclePos, isPossible := GetObstacleCoords(verticalPos, horizontalPos, currentDirection, input)
						if isPossible {
							_, isPlaced := placedObstacles[obstaclePos]
							if !isPlaced {
								placedObstacles[obstaclePos] = 0
								break
							}
						}
					}
				}
			} else if nextDir == 1 || nextDir == 3 {
				horizontalVisits := visitedCells[verticalPos]
				for _, el := range horizontalVisits {
					if slices.Contains(el, nextDir) {
						obstaclePos, isPossible := GetObstacleCoords(verticalPos, horizontalPos, currentDirection, input)
						if isPossible {
							_, isPlaced := placedObstacles[obstaclePos]
							if !isPlaced {
								placedObstacles[obstaclePos] = 0
							}
						}
					}
				}
			}

			verticalPos += directions[currentDirection].dirY
			horizontalPos += directions[currentDirection].dirX
			isPossible := IsCheckPossible(verticalPos, horizontalPos, len(input), len(input[0]))
			if !isPossible {
				isOver = true
				break
			}
			currentSymbol = input[verticalPos][horizontalPos]
		}
		if isOver {
			// We are out of the maze
			break
		}

		if currentSymbol == byte(wallSymbol) {
			verticalPos += -directions[currentDirection].dirY
			horizontalPos += -directions[currentDirection].dirX
			currentSymbol = input[verticalPos][horizontalPos]
			currentDirection = GetNextDir(currentDirection)
		}

	}
	fmt.Println("Result Part 2: ", len(placedObstacles))
}

// This does not cover some edge cases. Leaving for reference

/*
	func Part2(input []string, startX int, startY int) {
		var visitedCells map[string][]int = make(map[string][]int)
		var placedObstacles map[string]byte = make(map[string]byte)
		directions = make(map[int]Direction)
		directions[0] = Direction{dirX: 0, dirY: -1}
		directions[1] = Direction{dirX: 1, dirY: 0}
		directions[2] = Direction{dirX: 0, dirY: 1}
		directions[3] = Direction{dirX: -1, dirY: 0}

		wallSymbol := '#'

		isOver := false
		currentDirection := 0
		verticalPos := startX
		horizontalPos := startY
		currentSymbol := input[verticalPos][horizontalPos]
		for !isOver {
			for currentSymbol != byte(wallSymbol) {
				currentPosToKey := PositionToKey(verticalPos, horizontalPos)
				el, ok := visitedCells[currentPosToKey]
				// if we already visited the cell we can attempt placing and obstacle next to it
				if ok && slices.Contains(el, GetNextDir(currentDirection)) {
					obstaclePos, isPossible := GetObstacleCoords(verticalPos, horizontalPos, currentDirection, input)
					if isPossible {
						_, isPlaced := placedObstacles[obstaclePos]
						if !isPlaced {
							placedObstacles[obstaclePos] = 0
							// Reset and start again
							verticalPos = startX
							horizontalPos = startY
							currentDirection = 0
							currentSymbol = input[verticalPos][horizontalPos]
							break
						}
					}
				} else if !ok {
					visitedCells[currentPosToKey] = make([]int, 1)
					visitedCells[currentPosToKey][0] = currentDirection
				} else {
					if !slices.Contains(el, currentDirection) {
						visitedCells[currentPosToKey] = append(el, currentDirection)
					}

				}
				verticalPos += directions[currentDirection].dirY
				horizontalPos += directions[currentDirection].dirX
				isPossible := IsCheckPossible(verticalPos, horizontalPos, len(input), len(input[0]))
				if !isPossible {
					isOver = true
					break
				}
				currentSymbol = input[verticalPos][horizontalPos]
			}
			if isOver {
				// We are out of the maze
				break
			}

			if currentSymbol == byte(wallSymbol) {
				verticalPos += -directions[currentDirection].dirY
				horizontalPos += -directions[currentDirection].dirX
				currentSymbol = input[verticalPos][horizontalPos]
				currentDirection = GetNextDir(currentDirection)
			}

		}
		fmt.Println("Result Part 2: ", len(placedObstacles))
	}
*/

func GetObstacleCoords(verticalPos int, horizontalPos int, currentDirection int, input []string) (string, bool) {
	verticalObstaclePos := verticalPos + directions[currentDirection].dirY
	horizontalObstaclePos := horizontalPos + directions[currentDirection].dirX

	isPossible := IsCheckPossible(horizontalObstaclePos, verticalObstaclePos, len(input), len(input[0]))

	if !isPossible {
		return "", false
	}

	if input[verticalObstaclePos][horizontalObstaclePos] == byte('#') {
		return "", false
	}

	return PositionToKey(verticalObstaclePos, horizontalObstaclePos), true
}

func IsCheckPossible(x int, y int, height int, width int) bool {
	return (x < width) && (x >= 0) &&
		(y < height) && (y >= 0)
}

func GetNextDir(currentDir int) int {
	if currentDir == 3 {
		return 0
	} else {
		return currentDir + 1
	}
}

func PositionToKey(currentX int, currentY int) string {
	return strconv.Itoa(currentX) + ";" + strconv.Itoa(currentY)
}

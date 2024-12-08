package main

import (
	"adventofcode/shared"
	"log"
	"slices"
)

type Coordinates struct {
	row int
	col int
}

func (coord Coordinates) GetVectorToNext(nextCoords Coordinates) Coordinates {
	return Coordinates{row: coord.row - nextCoords.row, col: coord.col - nextCoords.col}
}

func GetAntinodeCoords(antennaCoords Coordinates, deltaVector Coordinates) Coordinates {
	return Coordinates{row: antennaCoords.row + deltaVector.row, col: antennaCoords.col + deltaVector.col}
}

func main() {
	lines, err := shared.ReadLines("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input %v", err)
	}
	Part1(lines)
	Part2(lines)

}

func Part1(lines []string) {
	knownAntennas := make(map[rune][]Coordinates)
	mapWidth := len(lines[0])
	mapHeight := len(lines)

	for i, line := range lines {
		for j, slot := range line {
			if slot == '.' {
				continue
			}
			_, ok := knownAntennas[slot]
			if !ok {
				knownAntennas[slot] = make([]Coordinates, 0)
			}
			knownAntennas[slot] = append(knownAntennas[slot], Coordinates{row: i, col: j})
		}
	}

	placedAntinodes := make([]Coordinates, 0)
	for _, sameAntennas := range knownAntennas {
		for i, uniqueAntenna := range sameAntennas {
			for j, neighbourAntenna := range sameAntennas {
				if i == j {
					continue
				}
				vectToNext := uniqueAntenna.GetVectorToNext(neighbourAntenna)
				antinodeCoords := GetAntinodeCoords(uniqueAntenna, vectToNext)
				// Checks for corner cases when not to add
				if (antinodeCoords.row < 0 || antinodeCoords.row >= mapHeight) ||
					(antinodeCoords.col < 0 || antinodeCoords.col >= mapWidth) {
					continue
				}
				doesOverlap := slices.ContainsFunc(placedAntinodes, func(coord Coordinates) bool {
					return (coord.col == antinodeCoords.col) && (coord.row == antinodeCoords.row)
				})
				if doesOverlap {
					continue
				}
				// if lines[antinodeCoords.row][antinodeCoords.col] != '.' {
				// 	continue
				// }

				placedAntinodes = append(placedAntinodes, antinodeCoords)
			}
		}
	}
	shared.PrintResultsInt(1, len(placedAntinodes))
}

func Part2(lines []string) {
	knownAntennas := make(map[rune][]Coordinates)
	mapWidth := len(lines[0])
	mapHeight := len(lines)

	for i, line := range lines {
		for j, slot := range line {
			if slot == '.' {
				continue
			}
			_, ok := knownAntennas[slot]
			if !ok {
				knownAntennas[slot] = make([]Coordinates, 0)
			}
			knownAntennas[slot] = append(knownAntennas[slot], Coordinates{row: i, col: j})
		}
	}

	placedAntinodes := make([]Coordinates, 0)
	for _, sameAntennas := range knownAntennas {
		for i, uniqueAntenna := range sameAntennas {
			for j := i + 1; j < len(sameAntennas); j++ {
				vectToNext := uniqueAntenna.GetVectorToNext(sameAntennas[j])

				tempPos := uniqueAntenna
				for (tempPos.row >= 0 && tempPos.row < mapHeight) &&
					(tempPos.col >= 0 && tempPos.col < mapWidth) {
					tempPos.col = tempPos.col - vectToNext.col
					tempPos.row = tempPos.row - vectToNext.row
				}
				tempPos.col += vectToNext.col
				tempPos.row += vectToNext.row

				for (tempPos.row >= 0 && tempPos.row < mapHeight) &&
					(tempPos.col >= 0 && tempPos.col < mapWidth) {
					doesOverlap := slices.ContainsFunc(placedAntinodes, func(coord Coordinates) bool {
						return (coord.col == tempPos.col) && (coord.row == tempPos.row)
					})
					if !doesOverlap {
						placedAntinodes = append(placedAntinodes, tempPos)
					}
					tempPos.col += vectToNext.col
					tempPos.row += vectToNext.row
				}

			}
		}
	}
	shared.PrintResultsInt(2, len(placedAntinodes))
}

package main

import (
	"fmt"
	"os"
	"strings"
)

func part1() int {
	input, _ := os.ReadFile("input.txt")

	grid := make([][]string, 0)
	for _, chars := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		rowChars := []string{}
		for _, char := range chars {
			rowChars = append(rowChars, string(char))
		}
		grid = append(grid, rowChars)
	}

	directions := [8][2]int{
		{1, 0},   // right
		{-1, 0},  // left
		{0, 1},   // down
		{0, -1},  // up
		{1, -1},  // rightUp
		{1, 1},   // rightDown
		{-1, 1},  // leftDown
		{-1, -1}, // leftUp
	}

	wordCount := 0
	for y_idx, row := range grid {
		for x_idx := range len(row) {
			for _, dirPair := range directions {
				checkForWordInDirection(grid, x_idx, y_idx, 0, &dirPair, &wordCount)
			}
		}
	}

	return wordCount
}

func checkForWordInDirection(grid [][]string, xCoord, yCoord, wordIdx int, dir *[2]int, wordCount *int) {
	WORD := [4]string{"X", "M", "A", "S"}

	newXCoord := xCoord + dir[0]
	newYCoord := yCoord + dir[1]

	if grid[yCoord][xCoord] == WORD[wordIdx] {
		newWordIdx := wordIdx + 1
		if newWordIdx == len(WORD) {
			// printFoundWord(grid, xCoord, yCoord, dir)
			*wordCount++
			return
		}

		if checkIfOutOfBounds(newXCoord, newYCoord, len(grid[0]), len(grid)) {
			return
		}

		checkForWordInDirection(grid, newXCoord, newYCoord, newWordIdx, dir, wordCount)
	}
}

func checkIfOutOfBounds(x, y, max_x, max_y int) bool {
	return x < 0 || x >= max_x || y < 0 || y >= max_y
}

func printFoundWord(grid [][]string, x, y int, dir *[2]int) {
	wordCoords := map[[2]int]bool{}
	for idx := 3; idx >= 0; idx-- {
		newX := x - (idx * dir[0])
		newY := y - (idx * dir[1])
		wordCoords[[2]int{newX, newY}] = true
	}

	const colorRed = "\033[0;31m"
	const colorNone = "\033[0m"

	for y_idx, row := range grid {
		for x_idx := range len(row) {
			if _, ok := wordCoords[[2]int{x_idx, y_idx}]; ok {
				fmt.Fprintf(os.Stdout, "\033[0;31m %s", grid[y_idx][x_idx])
			} else {
				fmt.Fprintf(os.Stdout, "\033[0m %s", grid[y_idx][x_idx])
			}
		}
		fmt.Println("")
	}
	fmt.Println()
}

func part2() int {
	input, _ := os.ReadFile("input.txt")

	grid := make([][]string, 0)
	for _, chars := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		rowChars := []string{}
		for _, char := range chars {
			rowChars = append(rowChars, string(char))
		}
		grid = append(grid, rowChars)
	}

	wordCount := 0
	for y_idx, row := range grid {
		for x_idx := range len(row) {
			if checkForXMAS(grid, x_idx, y_idx) {
				wordCount = wordCount + 1
			}
		}
	}

	return wordCount
}

func checkForXMAS(grid [][]string, x, y int) bool {
	if grid[y][x] != "A" {
		return false
	}
	// "A" is always in the middle just check for the four possible arrangements
	//
	// M.S or S.S or S.M or M.M
	// .A.    .A.    .A.    .A.
	// M.S    M.M    S.M    S.S
	possibleMappings := [4]map[[2]int]string{
		{
			{-1, -1}: "M", // top left
			{1, -1}:  "S", // top right
			{-1, 1}:  "M", // bottom left
			{1, 1}:   "S", // bottom right
		},
		{
			{-1, -1}: "S", // top left
			{1, -1}:  "S", // top right
			{-1, 1}:  "M", // bottom left
			{1, 1}:   "M", // bottom right
		},
		{
			{-1, -1}: "S", // top left
			{1, -1}:  "M", // top right
			{-1, 1}:  "S", // bottom left
			{1, 1}:   "M", // bottom right
		},
		{
			{-1, -1}: "M", // top left
			{1, -1}:  "M", // top right
			{-1, 1}:  "S", // bottom left
			{1, 1}:   "S", // bottom right
		},
	}

	for _, possibleMapping := range possibleMappings {
		mappingRightCount := 0
		for key, val := range possibleMapping {
			newX := x + key[0]
			newY := y + key[1]
			if checkIfOutOfBounds(newX, newY, len(grid[0]), len(grid)) {
				break
			}

			if grid[newY][newX] != val {
				break
			}

			mappingRightCount++
		}

		if mappingRightCount == 4 {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

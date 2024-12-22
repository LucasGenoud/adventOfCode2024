package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Pair struct {
	x, y int
}

func readInput() [][]string {
	var input [][]string

	readFile, err := os.Open("day4/input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	currentLine := 0
	for fileScanner.Scan() {
		currentLineValue := make([]string, 0)
		line := fileScanner.Text()
		for _, char := range line {
			currentLineValue = append(currentLineValue, string(char))
		}
		currentLine++
		input = append(input, currentLineValue)
	}
	readFile.Close()

	return input

}
func isXmasStar1(input [][]string, x int, y int, directions []Pair, target string) int {
	count := 0
	for _, direction := range directions {
		currentX := x
		currentY := y
		for _, targetChar := range target {
			if currentX < 0 || currentX >= len(input) {
				break
			}
			if currentY < 0 || currentY >= len(input[x]) {
				break
			}
			if input[currentX][currentY] != string(targetChar) {
				break
			}
			if string(targetChar) == string(target[len(target)-1]) {
				count++
			}
			currentX += direction.x
			currentY += direction.y
		}

	}
	return count
}

func star1() {
	target := "XMAS"

	total := 0
	directions := []Pair{Pair{1, 0}, Pair{-1, 0}, Pair{0, 1}, Pair{0, -1},
		Pair{1, 1}, Pair{-1, 1}, Pair{1, -1}, Pair{-1, -1}}
	input := readInput()
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			total += isXmasStar1(input, i, j, directions, target)

		}
	}

	fmt.Println(total)

}

func isXmasStar2(input [][]string, x int, y int, diagonals [][]Pair) bool {
	if input[x][y] != string('A') {
		return false
	}

	for _, diagonal := range diagonals {
		relevantChars := []string{"M", "S"}

		for _, direction := range diagonal {
			currentX := x + direction.x
			currentY := y + direction.y
			if currentX < 0 || currentX >= len(input) {
				return false
			}
			if currentY < 0 || currentY >= len(input[x]) {
				return false
			}
			currentChar := input[currentX][currentY]

			index := slices.Index(relevantChars, currentChar)
			if index == -1 {
				return false
			}
			relevantChars = slices.Delete(relevantChars, index, index+1)

		}
		if len(relevantChars) != 0 {
			return false
		}
	}
	return true

}

func star2() {

	total := 0
	directions1 := []Pair{Pair{1, 1}, Pair{-1, -1}}
	directions2 := []Pair{Pair{1, -1}, Pair{-1, 1}}
	directions := [][]Pair{directions1, directions2}
	input := readInput()
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if isXmasStar2(input, i, j, directions) {
				total++

			}

		}
	}

	fmt.Println(total)
}

func main() {
	star1()
	star2()

}

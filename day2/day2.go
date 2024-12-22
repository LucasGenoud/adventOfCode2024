package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getArray() [][]int {
	var list [][]int
	readFile, err := os.Open("day2/input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	listParsedInt := make([]int, 0)

	for fileScanner.Scan() {
		values := strings.Split(fileScanner.Text(), " ")
		for _, value := range values {
			valueInt, _ := strconv.Atoi(value)
			listParsedInt = append(listParsedInt, valueInt)
		}
		list = append(list, listParsedInt)
		listParsedInt = make([]int, 0)
	}
	readFile.Close()

	return list
}
func goesUp(step1 int, step2 int) bool {
	return step1 > step2
}
func isLevelSafe(level []int) bool {
	firstStepGoesUp := goesUp(level[0], level[1])
	for i := 0; i < len(level)-1; i++ {
		step1 := level[i]
		step2 := level[i+1]
		delta := math.Abs(float64(step1 - step2))
		if delta > 3 {
			return false
		}
		if delta == 0 {
			return false
		}
		if goesUp(step1, step2) != firstStepGoesUp {
			return false
		}
	}
	return true
}
func remove(slice []int, s int) []int {
	newArray := make([]int, 0)
	for i := 0; i < len(slice); i++ {
		if i == s {
			continue
		} else {
			newArray = append(newArray, slice[i])
		}
	}
	return newArray
}

func star1() {
	safeLevels := 0
	levels := getArray()
	for _, level := range levels {
		if isLevelSafe(level) {
			safeLevels++
		}
	}
	fmt.Println(safeLevels)
}
func star2() {
	safeLevels := 0
	levels := getArray()
	for _, level := range levels {
		if isLevelSafe(level) {
			safeLevels++
		} else {
			for step := 0; step < len(level); step++ {
				newLevel := remove(level, step)
				if isLevelSafe(newLevel) {
					safeLevels++
					break
				}
			}
		}
	}
	fmt.Println(safeLevels)
}

func main() {
	star2()

}

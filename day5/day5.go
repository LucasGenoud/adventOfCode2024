package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	x, y int
}

func convertToInt(s string) int {
	number, _ := strconv.Atoi(s)
	return number
}

func readInput() ([]Pair, [][]int) {
	var rules []Pair
	var sequences [][]int
	readFile, err := os.Open("day5/input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.Contains(line, "|") {
			split := strings.Split(line, "|")
			newRule := Pair{x: convertToInt(split[0]), y: convertToInt(split[1])}
			rules = append(rules, newRule)
		} else {
			split := strings.Split(line, ",")
			var parsedList []int
			for _, s := range split {
				parsedList = append(parsedList, convertToInt(s))
			}
			sequences = append(sequences, parsedList)
		}
	}
	readFile.Close()

	return rules, sequences

}

func star1() {
	rules, sequences := readInput()
	print(rules)
	print(sequences)
}

func star2() {

}
func main() {
	star1()
	star2()

}

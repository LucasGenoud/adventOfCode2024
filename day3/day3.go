package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getString() string {
	fileContent, err := os.ReadFile("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(fileContent)
}

func star1() {
	regexExpression := regexp.MustCompile(`mul\(\d+,\d+\)`)
	fileContent := getString()
	expressions := regexExpression.FindAllString(fileContent, -1)
	total := 0
	for _, expression := range expressions {
		value := strings.Replace(expression, "mul(", "", -1)
		value = strings.Replace(value, ")", "", -1)
		values := strings.Split(value, ",")
		value1, _ := strconv.Atoi(values[0])
		value2, _ := strconv.Atoi(values[1])
		total += value1 * value2
	}
	print(total)

}
func star2() {
	regexExpression := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	fileContent := getString()
	expressions := regexExpression.FindAllString(fileContent, -1)
	total := 0
	processMul := true
	for _, expression := range expressions {
		if expression == "do()" {
			processMul = true
			continue
		}
		if expression == "don't()" {
			processMul = false
			continue
		}
		if !processMul {
			continue
		}
		value := strings.Replace(expression, "mul(", "", -1)
		value = strings.Replace(value, ")", "", -1)
		values := strings.Split(value, ",")
		value1, _ := strconv.Atoi(values[0])
		value2, _ := strconv.Atoi(values[1])
		total += value1 * value2
	}
	print(total)

}

func main() {
	star1()

}

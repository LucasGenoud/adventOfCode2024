package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	star2()
}
func getSortedInts() ([]int, []int) {
	var list1 []int
	var list2 []int
	readFile, err := os.Open("day1/input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		values := strings.Split(fileScanner.Text(), "   ")
		value1Int, _ := strconv.Atoi(values[0])
		value2Int, _ := strconv.Atoi(values[1])
		list1 = append(list1, value1Int)
		list2 = append(list2, value2Int)
	}
	sort.Ints(list1)
	sort.Ints(list2)
	readFile.Close()

	return list1, list2
}
func star1() {
	list1, list2 := getSortedInts()
	totalDiff := 0

	for i := 0; i < len(list1); i++ {
		diff := list1[i] - list2[i]
		if diff < 0 {
			diff = diff * -1
		}
		totalDiff += diff
	}
	fmt.Println(totalDiff)

}

func star2() {
	list1, list2 := getSortedInts()
	total := 0

	for i := 0; i < len(list1); i++ {
		count := countIntsInSortedList(list2, list1[i])
		total += count * list1[i]
	}
	fmt.Println(total)
}

func countIntsInSortedList(list []int, value int) int {
	count := 0
	for i := 0; i < len(list); i++ {
		if list[i] > value {
			return count
		}
		if list[i] == value {
			count++
		}

	}
	return count

}

package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func part1(data string) int {
	var leftList []int
	var rightList []int

	lines := strings.Split(data, "\n")
	for _, line := range lines {
		words := strings.Fields(line)
		left, _ := strconv.Atoi(words[0])
		right, _ := strconv.Atoi(words[1])
		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	sort.Sort(sort.IntSlice(leftList))
	sort.Sort(sort.IntSlice(rightList))

	total := 0
	for i := 0; i < len(leftList); i++ {
		total += abs(leftList[i] - rightList[i])
	}

	return total
}

func part2(data string) int {
	var leftList []int
	repeats := make(map[int]int)

	lines := strings.Split(data, "\n")
	for _, line := range lines {
		words := strings.Fields(line)
		left, _ := strconv.Atoi(words[0])
		right, _ := strconv.Atoi(words[1])
		leftList = append(leftList, left)

		repeats[right]++
	}

	total := 0
	for _, i := range leftList {
		total += i * repeats[i]
	}

	return total
}

func main() {
	fmt.Println("Day 01")

	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Part 1: ", part1(string(file)))
	fmt.Println("Part 2: ", part2(string(file)))
}

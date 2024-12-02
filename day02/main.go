package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func checkSafe(report []int) bool {
	increasing := report[0] <= report[1]

	for i := 1; i < len(report); i++ {
		diff := int(math.Abs(float64(report[i-1] - report[i])))
		if (increasing && report[i-1] > report[i]) ||
			(!increasing && report[i-1] < report[i]) ||
			(diff < 1 || diff > 3) {
			return false
		}
	}

	return true
}

func part1(reports [][]int) int {
	total := 0
	for _, report := range reports {
		if checkSafe(report) {
			total += 1
		}
	}
	return total
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func part2(reports [][]int) int {
	total := 0
	for _, report := range reports {
		if checkSafe(report) {
			total += 1
		} else {
			for i := 0; i < len(report); i++ {
				newSlice := make([]int, len(report))
				copy(newSlice, report)
				if checkSafe(append(newSlice[:i], newSlice[i+1:]...)) {
					total += 1
					break
				}
			}
		}
	}
	return total
}

func main() {
	fmt.Println("Day 02")

	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.Split(string(file), "\n")

	reports := make([][]int, 0)
	for _, line := range lines {
		report := make([]int, 0)
		for _, f := range strings.Fields(line) {
			num, _ := strconv.Atoi(f)
			report = append(report, num)
		}
		reports = append(reports, report)
	}

	fmt.Println("Part 1: ", part1(reports))
	fmt.Println("Part 2: ", part2(reports))
}

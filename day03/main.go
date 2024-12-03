package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1(data string) int {
	r, _ := regexp.Compile("mul\\(\\d{1,3},\\d{1,3}\\)")
	replace, _ := regexp.Compile("[^\\d,]")
	matchs := r.FindAllString(data, -1)

	total := 0
	for _, match := range matchs {
		numbers := strings.Split(replace.ReplaceAllString(match, ""), ",")
		n1, _ := strconv.Atoi(numbers[0])
		n2, _ := strconv.Atoi(numbers[1])
		total += n1 * n2
	}

	return total
}

func part2(data string) int {
	r, _ := regexp.Compile("(mul\\(\\d{1,3},\\d{1,3}\\))|(do(n't)*\\(\\))")
	replace, _ := regexp.Compile("[^\\d,]")
	matchs := r.FindAllString(data, -1)

	total := 0
	enabled := true
	for _, match := range matchs {
		if match == "don't()" {
			enabled = false
		} else if match == "do()" {
			enabled = true
		} else if enabled {
			numbers := strings.Split(replace.ReplaceAllString(match, ""), ",")
			n1, _ := strconv.Atoi(numbers[0])
			n2, _ := strconv.Atoi(numbers[1])
			total += n1 * n2
		}
	}

	return total
}

func main() {
	fmt.Println("Day 03")

	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	data := string(file)

	fmt.Println("Part 1: ", part1(data))
	fmt.Println("Part 2: ", part2(data))
}

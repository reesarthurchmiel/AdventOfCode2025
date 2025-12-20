package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputString string

func main() {
	//day3_part1()
	day3_part2()
}

func day3_part1() {
	result := CalculateTotalJoltage(inputString)
	fmt.Println(result)
}

func CalculateTotalJoltage(input string) int {
	lines := strings.Split(input, "\n")

	total := 0
	for _, line := range lines {
		line = strings.TrimRight(line, "\r")
		nums := make([]int, len(line))
		for i, v := range strings.Split(line, "") {
			asInt, _ := strconv.Atoi(v)
			nums[i] = asInt
		}

		joltage := CalculateJoltage(nums)

		total += joltage
	}

	return total
}

func CalculateJoltage(nums []int) int {
	max := 0
	maxI := -1

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > max {
			max = nums[i]
			maxI = i
		}
	}

	maxSecond := 0
	for i := maxI + 1; i < len(nums); i++ {
		if nums[i] > maxSecond {
			maxSecond = nums[i]
		}
	}
	return max*10 + maxSecond
}

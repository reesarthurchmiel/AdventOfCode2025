package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day3_part2() {
	result := CalculateTotal(inputString)
	fmt.Println(result)
}

type MemoKey struct {
	start     int
	leftToGet int
}

func CalculateTotal(input string) int {
	lines := strings.Split(input, "\n")

	total := 0
	for _, line := range lines {
		line = strings.TrimRight(line, "\r")
		nums := make([]int, len(line))
		for i, v := range strings.Split(line, "") {
			asInt, _ := strconv.Atoi(v)
			nums[i] = asInt
		}

		memo := make(map[MemoKey]int)
		joltage := CalculateBigJoltage(nums, 0, 12, memo)

		total += joltage
	}

	return total
}

func CalculateBigJoltage(nums []int, start int, leftToGet int, memo map[MemoKey]int) int {
	if leftToGet == 0 {
		return 0
	}

	memoKey := MemoKey{start, leftToGet}
	if val, ok := memo[memoKey]; ok {
		return val
	}

	if leftToGet == len(nums)-start {
		result := 0
		for i := start; i < len(nums); i++ {
			result *= 10
			result += nums[i]
		}

		memo[memoKey] = result
		return result
	}

	if leftToGet > len(nums)-start {
		panic("Not got enough left in nums for how many we need")
	}

	with := (nums[start] * IntPow(10, leftToGet-1)) + CalculateBigJoltage(nums, start+1, leftToGet-1, memo)
	without := CalculateBigJoltage(nums, start+1, leftToGet, memo)

	result := max(with, without)
	memo[memoKey] = result

	return result
}

func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}

	if m == 1 {
		return n
	}

	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

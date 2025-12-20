package main

import (
	_ "embed"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputString string

func main() {
	lines := GetLines(inputString)
	freshes := GetRanges(lines)

	result := part1(lines, freshes)
	fmt.Println("Part 1:", result)

	result2 := part2(freshes)
	fmt.Println("Part 2:", result2)
}

func part1(lines []string, freshes []FreshRange) int {
	maxI := 0

	for i, line := range lines {
		maxI = i
		if line == "" {
			break
		}
	}

	result := 0

	for i := maxI + 1; i < len(lines); i++ {
		id, _ := strconv.Atoi(lines[i])

		for _, fresh := range freshes {
			if fresh.left <= id && id <= fresh.right {
				result++
				break
			}
		}
	}

	return result
}

func part2(freshes []FreshRange) int {
	result := 0

	for _, fresh := range freshes {
		result += (fresh.right - fresh.left) + 1 // +1 because the range is inclusive of both ends
	}

	return result
}

// inclusive
type FreshRange struct {
	left  int
	right int
}

func GetRanges(lines []string) []FreshRange {
	var freshes []FreshRange

	for _, line := range lines {
		if line == "" {
			break
		}

		leftStr := strings.Split(line, "-")[0]
		rightStr := strings.Split(line, "-")[1]

		left, _ := strconv.Atoi(leftStr)
		right, _ := strconv.Atoi(rightStr)

		freshes = append(freshes, FreshRange{left, right})
	}

	sort.Slice(freshes, func(i, j int) bool {
		if freshes[i].left == freshes[j].left {
			return freshes[i].right < freshes[j].right
		}
		return freshes[i].left < freshes[j].left
	})

	for i := 0; i < len(freshes)-1; {
		if freshes[i].right >= freshes[i+1].left {
			// merge
			freshes[i+1] = FreshRange{freshes[i].left, max(freshes[i].right, freshes[i+1].right)}
			freshes = slices.Delete(freshes, i, i+1)
		} else {
			i++
		}
	}

	return freshes
}

func GetLines(input string) []string {
	lines := strings.Split(input, "\n")
	for i, v := range lines {
		lines[i] = strings.TrimRight(v, "\r")
	}
	return lines
}

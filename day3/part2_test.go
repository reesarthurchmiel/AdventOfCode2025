package main

import "testing"

func TestCalculateBigJoltage(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1}, 98},
		{[]int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9}, 89},
		{[]int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8}, 78},
		{[]int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1}, 92},
	}

	for _, tt := range tests {
		result := CalculateBigJoltage(tt.input, 0, 2)
		if result != tt.want {
			t.Errorf("Even the line by line is fucked, expected %d got %d", tt.want, result)
		}
	}
}

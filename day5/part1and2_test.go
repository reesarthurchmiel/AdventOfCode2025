package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetRanges(t *testing.T) {
	input := []string{"1-4", "3-8", "9-11", "4-6"}

	result := GetRanges(input)

	for _, rang := range result {
		fmt.Printf("%d-%d\n", rang.left, rang.right)
	}

	if !reflect.DeepEqual(result, []FreshRange{{1, 8}, {9, 11}}) {
		t.Fail()
	}
}

package main

/*
import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputString string

func main() {
	var idRanges []string = strings.Split(inputString, ",")
	var idsToAdd []int64

	for _, idRange := range idRanges {
		idsToAdd = checkIdRange(idRange, idsToAdd)
	}

	answer := int64(0)
	for _, v := range idsToAdd {
		answer += v
	}

	fmt.Println(answer)
}

func checkIdRange(idRange string, idsToAdd []int64) []int64 {
	ends := strings.Split(idRange, "-")
	left, _ := strconv.ParseInt(ends[0], 0, 0)
	right, _ := strconv.ParseInt(ends[1], 0, 0)

	for id := left; id <= right; id++ {
		idString := strconv.Itoa(int(id))
		if len(idString)%2 == 1 {
			continue
		}

		leftSide := idString[:len(idString)/2]
		rightSide := idString[len(idString)/2:]

		if leftSide == rightSide {
			idsToAdd = append(idsToAdd, id)
		}
	}
	return idsToAdd
} */

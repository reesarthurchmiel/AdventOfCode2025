package main

import (
	_ "embed"
	"fmt"
	"github.com/dlclark/regexp2"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputString string
var idRegex *regexp2.Regexp

func main() {
	idRegex = regexp2.MustCompile(`^(\d+)\1{1,}$`, regexp2.None)

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

		if ok, _ := idRegex.MatchString(idString); ok {
			idsToAdd = append(idsToAdd, id)
		}
	}
	return idsToAdd
}

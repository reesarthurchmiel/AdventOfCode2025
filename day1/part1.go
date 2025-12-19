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
	var lines []string = strings.Split(inputString, "\r\n")
	answer := 0
	dial := 50
	for _, line := range lines {
		if len(line) == 0 {
			continue // empty line
		}
		i, _ := strconv.ParseInt(line[1:], 10, 0)

		movement := int(i)
		if line[0] == 'L' {
			dial -= movement
		} else {
			dial += movement
		}

		dial = dial % 100
		if dial == 0 {
			answer++
		}
	}

	fmt.Println("Answer is:", answer)
}
*/

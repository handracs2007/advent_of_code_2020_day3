package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func countTrees(lines []string, right int, down int) int {
	cols := len(lines[0])
	rows := len(lines)
	curCol, curRow := 0, 0
	tCount := 0

	for curRow < rows {
		curCol += right
		curCol %= cols

		curRow += down

		if curRow < rows {
			if rune(lines[curRow][curCol]) == '#' {
				tCount++
			}
		}
	}

	return tCount
}

func main() {
	fc, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Println(err)
		return
	}

	lines := strings.Split(string(fc), "\n")

	count := 1
	count *= countTrees(lines, 1, 1)
	count *= countTrees(lines, 3, 1)
	count *= countTrees(lines, 5, 1)
	count *= countTrees(lines, 7, 1)
	count *= countTrees(lines, 1, 2)

	fmt.Println(count)
}

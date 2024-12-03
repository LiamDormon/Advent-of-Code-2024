package solutions

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"utils"
)

func Part2() int {
	leftList := []int{}
	rightList := []int{}

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		halves := strings.Split(line, "   ")

		left, err := strconv.Atoi(halves[0])
		check(err)
		right, err := strconv.Atoi(halves[1])
		check(err)

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	similarityScore := 0

	for _, left := range leftList {
		count := utils.CountOccurrences(rightList, left)

		similarityScore += left * count
	}

	return similarityScore
}
